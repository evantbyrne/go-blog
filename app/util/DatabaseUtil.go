package util

import "database/sql"
import "log"

import _ "github.com/lib/pq"

func DatabaseConnect() *sql.DB {
	db, err := sql.Open("postgres", "user=evantbyrne dbname=go_blog sslmode=disable")
	if (err != nil) {
		log.Fatal(err)
	}

	return db
}