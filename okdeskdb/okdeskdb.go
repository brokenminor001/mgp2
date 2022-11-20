package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// "user=bi_soyuzintegro password=30a08bc2d9d782eb host=188.42.59.228 port=10100 dbname=bi_db_soyuzintegro sslmode=disable"
const (
	host     = "188.42.59.228"
	port     = 10100
	user     = "bi_soyuzintegro"
	password = "30a08bc2d9d782eb"
	dbname   = "bi_db_soyuzintegro"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var ticket_id string
	connectionstring := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select sequential_id from issues order by id desc limit 1;")
	CheckError(err)
	for rows.Next() {

		err = rows.Scan(&ticket_id)
		CheckError(err)

	}
	fmt.Print(ticket_id)
}
