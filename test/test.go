package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "172.17.0.2"
	port     = 5432
	user     = "asmolin"
	password = "oadnpvia"
	dbname   = "megapolis"
)

func InsertnewticketID(id string, ticket_id string) {
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("update tickets set okdesk_id=$1 where ticket_id=$2", id, ticket_id)
	if err != nil {
		panic(err)
	}
	log.Print(result)

}

func main() {
	s1 := "5555"
	s2 := "ГКМ15723244"
	//s2 := "ГКМ00001"
	InsertnewticketID(s1, s2)
}
