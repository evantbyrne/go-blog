package util

import "database/sql"
import "log"

import "github.com/coopernurse/gorp"
import _ "github.com/lib/pq"

import "../model/dto"

func Database() *sql.DB {
	db, err := sql.Open("postgres", "user=evantbyrne dbname=go_blog sslmode=disable")
	if (err != nil) {
		log.Fatal(err)
	}

	return db
}

func DatabaseMap() *gorp.DbMap {
	db := Database()
	dbmap := &gorp.DbMap{ Db: db, Dialect: gorp.PostgresDialect{} }
	dbmap.AddTableWithName(dto.Article{}, "article").SetKeys(true, "Id")
	return dbmap
}