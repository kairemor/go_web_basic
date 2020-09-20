package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

var users []User

func init() {
	fmt.Println("Go & DB [postgresql]")

	db, err = sql.Open("postgres", "dbname=node_api host=localhost user=postgres password=postgre")

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id ,"firstName", "email" from "Users"`)

	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id, firstName, email string
		rows.Scan(&id, &firstName, &email)

		u := User{ID: id, Name: firstName, Email: email}

		users = append(users, u)

		fmt.Printf("%s %s %s \n", id, firstName, email)
	}

	if rows.Err() != nil {
		log.Fatal(err)
	}
}
