package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Go & DB [postgresql]")

	db, err := sql.Open("postgres", "dbname=node_api host=localhost user=postgres password=postgre")

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id ,"firstName", "lastName" from "Users"`)

	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id, firstName, lastName string
		rows.Scan(&id, &firstName, &lastName)

		fmt.Printf("%s %s %s \n", id, firstName, lastName)
	}

	if rows.Err() != nil {
		log.Fatal(err)
	}
}
