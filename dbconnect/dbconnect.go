package dbconnect

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "172.17.0.3"
	port     = 5432
	user     = "asmolin"
	password = "oadnpvia"
	dbname   = "megapolis"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Getticketid() string {
	var ticket_id string
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select ticket_id from tickets order by id desc limit 1;")
	CheckError(err)
	for rows.Next() {

		err = rows.Scan(&ticket_id)
		CheckError(err)

	}
	return ticket_id
}

func Insertnewticket(tick_id string) {
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into tickets (ticket_id,status) values ($1,$2)",
		tick_id, 0)
	if err != nil {
		panic(err)
	}
	log.Print(result)

}

func InsertnewticketID(id string, ticket_id string) {
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update tickets set okdesk_id=$1 where ticket_id=$2",
		id, ticket_id)
	if err != nil {
		panic(err)
	}
	log.Print(result)

}
func SelectTicketById(id string) string {
	var ticket_id string
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select ticket_id from tickets where okdesk_id=$1", id)
	CheckError(err)
	for rows.Next() {

		err = rows.Scan(&ticket_id)
		CheckError(err)

	}

	return ticket_id

}

func GetStatusID(status string) string {
	var statuschek string
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select status from tickets where okdesk_id=$1", status)
	CheckError(err)
	for rows.Next() {

		err = rows.Scan(&statuschek)
		CheckError(err)

	}

	return statuschek

}
func UpdateStatusOne(okdeskid string) {
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update tickets set status='1' where okdesk_id=$1",
		okdeskid)
	if err != nil {
		panic(err)
	}
	log.Print(result)
}

func UpdateUpdate(ticketid string) {
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update tickets set update=$1 where ticket_id=$2",
		ticketid)
	if err != nil {
		panic(err)
	}
	log.Print(result)
}
func UpdateChek(tick_id string) string {
	var updcheck string
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select update from tickets where okdesk_id=$1", tick_id)
	CheckError(err)
	for rows.Next() {

		err = rows.Scan(&updcheck)
		CheckError(err)

	}

	return updcheck

}
func GetOkdeskID(tic_id string) string {
	var okdeskid string
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select ticket_id from tickets where okdesk_id=$1", tic_id)
	CheckError(err)
	for rows.Next() {

		err = rows.Scan(&okdeskid)
		CheckError(err)

	}

	return okdeskid

}
