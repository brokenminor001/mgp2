package sendmail

import (
	"gopkg.in/gomail.v2"
)

func Sendassigned(ticket string) {

	m := gomail.NewMessage()
	m.SetHeader("From", "megapolis@soyuzintegro.ru")
	m.SetHeader("To", "operu_it_inbox_robot@gkm.ru")
	m.SetHeader("Subject", "Incident/"+ticket+"/Assigned\r\n\r\n")
	m.SetBody("text/plain", "<incident><number>"+ticket+"</number><nameCharge>АутсорсСоюзинтегро</nameCharge><phoneCharge>+7 (900) 500-00-00</phoneCharge></incident>\r\n")

	// Send the email to Bob
	d := gomail.NewPlainDialer("mail.soyuzintegro.ru", 25, "asmolin@soyuzintegro.ru", "Oadnpvia04!!")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

func SendStartWork(ticket string) {

	m := gomail.NewMessage()
	m.SetHeader("From", "megapolis@soyuzintegro.ru")
	m.SetHeader("To", "operu_it_inbox_robot@gkm.ru")
	m.SetHeader("Subject", "Incident/"+ticket+"/Update\r\n\r\n")
	m.SetBody("text/plain", `<?xml version="1.0" encoding="UTF-8"?><incident><number>`+ticket+`</number><incidentId>691458</incidentId><status>В работе</status></incident>`)

	// Send the email to Bob
	d := gomail.NewPlainDialer("mail.soyuzintegro.ru", 25, "asmolin@soyuzintegro.ru", "Oadnpvia04!!")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

func SendEndWork(ticket string) {

	m := gomail.NewMessage()
	m.SetHeader("From", "megapolis@soyuzintegro.ru")
	m.SetHeader("To", "operu_it_inbox_robot@gkm.ru")
	m.SetHeader("Subject", "Incident/"+ticket+"/Update\r\n\r\n")
	m.SetBody("text/palin", `<?xml version="1.0" encoding="UTF-8"?><incident><number>`+ticket+`</number><comment>[title:Администратор СоюзИнтегро]: Задание 682406/480091 мастер отчитался о завершении работ</comment></incident>`)

	// Send the email to Bob
	d := gomail.NewPlainDialer("mail.soyuzintegro.ru", 25, "asmolin@soyuzintegro.ru", "Oadnpvia04!!")
	if err := d.DialAndSend(m); err != nil {
		panic(err)

	}

}
