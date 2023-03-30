package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/brokenminor001/mgp2/dbconnect"
	"github.com/brokenminor001/mgp2/mgp"
	"github.com/brokenminor001/mgp2/sendmail"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

const (
	newticket = "Incident/Create"
	update    = "update"
	endwork   = " Заявка перешла в статус Окончил работу"
	startwork = " Заявка перешла в статус Приступил к работе"
)

type Newpost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company_id  string `json:"company_id"`
}

type Updatepost struct {
	Content   string `json:"content`
	Author_id string `json:"content`
}

var text string
var Finaltext string
var FileName string
var Attach bool = false

func getmsg() {

	// Connect to server

	c, err := client.DialTLS("imap.yandex.ru:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login("brokenminor001@yandex.ru", "skgvdtxrafxjebxc"); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	// Get the last message
	if mbox.Messages == 0 {
		log.Fatal("No message in mailbox")
	}
	seqSet := new(imap.SeqSet)
	seqSet.AddNum(mbox.Messages)

	// Get the whole message body
	var section imap.BodySectionName
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 1)
	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			log.Fatal(err)
		}
	}()

	msg := <-messages
	if msg == nil {
		log.Fatal("Server didn't returned message")
	}

	r := msg.GetBody(&section)
	if r == nil {
		log.Fatal("Server didn't returned message body")
	}

	// Create a new mail reader
	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Fatal(err)
	}

	// Print some info about the message
	header := mr.Header
	// if date, err := header.Date(); err == nil {
	// 	log.Println("Date:", date)
	// }
	// if from, err := header.AddressList("From"); err == nil {
	// 	log.Println("From:", from)
	// }
	// if to, err := header.AddressList("To"); err == nil {
	// 	log.Println("To:", to)
	// }
	// if subject, err := header.Subject(); err == nil {
	// 	log.Println("Subject:", subject)

	// }
	subject, err := header.Subject()
	if err != nil {
		log.Fatal(err)
	}

	// Process each message's part
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			// This is the message's text (can be plain-text or HTML)
			b, _ := ioutil.ReadAll(p.Body)

			text = string(b)

		case *mail.AttachmentHeader:
			// This is an attachment
			filename, _ := h.Filename()

			output, err := os.Create(filename)
			if err != nil {
				fmt.Println("Error while creating", "-", err)
			}
			defer output.Close()
			//--------------

			//записываем байты в файл
			_, err = io.Copy(output, p.Body)
			if err != nil {
				fmt.Println("Error while downloading", "-", err)

			}
			Attach = true
			FileName = filename
		}

	}

	fmt.Print(FileName)
	fmt.Print(FileName)

	updstsone := strings.Split(subject, "/")
	updsts := updstsone[0]
	updtextone := strings.Split(text, ":")
	updtexttwo := updtextone[1]
	updticid := updtexttwo[:13]
	updtext := updtextone[2]

	// ///////////// main process /////////////
	if subject == newticket {
		fmt.Print("Check for new ticket\n")
		parts := strings.Split(text, ":")
		fmt.Print(parts)
		tickettextone := parts[1]
		fmt.Print(tickettextone)
		tickettexttwo := strings.Split(tickettextone, "Состояние")
		tickettext := tickettexttwo[0]
		tic := tickettext[:14]

		var ticketid string = dbconnect.Getticketid()

		var compare bool = ticketid != tic

		if compare == true {
			if Attach == false {
				dbconnect.Insertnewticket(tic)
				themetextone := parts[12]
				fmt.Print(themetextone)
				themetexttwo := strings.Split(themetextone, "Серийный номер")
				fmt.Print(themetexttwo)
				var themetext string = themetexttwo[0]
				serialone := parts[13]
				serialtwo := strings.Split(serialone, "Требуется диагностика и список ЗИП для возможного ремонта?")
				var finaltext string = themetext + "\n" + "Серийный номер:" + serialtwo[0]

				var post = Newpost{subject, finaltext, "22"}
				data, err := json.Marshal(post)
				fmt.Println(string(data))

				r := bytes.NewBuffer(data)
				resp, err := http.Post("https://soyuzintegro.okdesk.ru/api/v1/issues/?api_token=5683dfe9931d8e891acb4943a76bde8fc2edeada", "application/json", r)
				fmt.Printf("%v %v", err, resp)
				log.Print("no new tickets")
				sendmail.Sendassigned(ticketid)
				//time.Sleep(30 * time.Second)
				var id string = mgp.GetNewTicketID()
				dbconnect.InsertnewticketID(id, tic)
			}
			if Attach == true {
				dbconnect.Insertnewticket(tic)
				themetextone := parts[12]
				fmt.Print(themetextone)
				themetexttwo := strings.Split(themetextone, "Серийный номер")
				fmt.Print(themetexttwo)
				var themetext string = themetexttwo[0]
				serialone := parts[13]
				serialtwo := strings.Split(serialone, "Требуется диагностика и список ЗИП для возможного ремонта?")
				var finaltext string = themetext + "\n" + "Серийный номер:" + serialtwo[0]

				cmd := exec.Command("curl", "-H", "ontent-Type: multipart/form-data", "-F", "issue[title]="+subject, "-F", "issue[company_id]=22", "-F", "issue[description]="+finaltext, "-F", "issue[attachments][0][attachment]=@/go/"+FileName, "https://soyuzintegro.okdesk.ru/api/v1/issues/?api_token=5683dfe9931d8e891acb4943a76bde8fc2edeada")

				cmd.Run()

			}

		} else {
			fmt.Print("no new tickets")

		}

		if updsts == update {
			var checkupdate string = dbconnect.UpdateChek(updticid)

			if checkupdate != updtext {
				var id string = dbconnect.GetOkdeskID(updticid)
				var post = Updatepost{updtext, "6"}
				data, err := json.Marshal(post)
				fmt.Println(string(data))

				r := bytes.NewBuffer(data)
				resp, err := http.Post("https://soyuzintegro.okdesk.ru/api/v1/issues/"+id+"/comments?api_token=5683dfe9931d8e891acb4943a76bde8fc2edeada", "application/json", r)
				fmt.Printf("%v %v", err, resp)
			}
		} else {

			fmt.Print("Nothing")
		}

	}
}
func status() {
	fmt.Print("start function status")

	var status string = mgp.GetStatus()
	fmt.Print(status)
	fmt.Print(startwork)
	if status == startwork {
		fmt.Print("check")
		var okdesk_id string = mgp.GetNewTicketID()
		var tic_id string = dbconnect.SelectTicketById(okdesk_id)
		var statuscheck string = dbconnect.GetStatusID(okdesk_id)
		var check string = "0"
		if statuscheck == check {
			sendmail.SendStartWork(tic_id)
			dbconnect.UpdateStatusOne(okdesk_id)

		} else {
			fmt.Print("status chek not complite")
		}

	} else if status == endwork {
		fmt.Print("ENDWORK!!!!")
		var okdesk_id string = mgp.GetNewTicketID()
		var tic_id string = dbconnect.SelectTicketById(okdesk_id)
		var statuscheck string = dbconnect.GetStatusID(okdesk_id)
		var check string = "1"
		if statuscheck == check {
			sendmail.SendEndWork(tic_id)
			dbconnect.UpdateStatusTwo(okdesk_id)
		}
	} else {

		fmt.Print("No new status")

	}

}

func period() {
	for {
		getmsg()
		status()
		time.Sleep(120 * time.Second)
	}
}

func main() {

	period()

}
