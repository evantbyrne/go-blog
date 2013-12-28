package util

import "database/sql"
import "log"

import "github.com/coopernurse/gorp"
import _ "github.com/lib/pq"

import "../model/dto"

type Database struct {
	Connection *sql.DB
	Opened bool
}

func (db *Database) Open() *sql.DB {
	if !db.Opened {
		connection, err := sql.Open("postgres", "user=evantbyrne dbname=go_blog sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		db.Connection = connection
		db.Opened = true
	}

	return db.Connection
}

func (db *Database) Gorp() *gorp.DbMap {
	connection := db.Open()
	dbmap := &gorp.DbMap{ Db: connection, Dialect: gorp.PostgresDialect{} }
	dbmap.AddTableWithName(dto.Article{}, "article").SetKeys(true, "Id")
	return dbmap
}