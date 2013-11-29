package main

import "database/sql"
import "log"
import "net/http"

import _ "github.com/lib/pq"

import "./app/model/dto"
import "./app/util"

// --- Pages ---

func PageIndex(response http.ResponseWriter, request *http.Request) {
	// Index
	if (request.URL.Path == "/") {
		db, err := sql.Open("postgres", "user=evantbyrne dbname=go_blog sslmode=disable")
		if (err != nil) {
			log.Fatal(err)
		}

		res := make([]dto.Article, 0)
		var (
			id int
			title string
			content string
		)
		rows, err := db.Query("select * from article")
		if (err != nil) {
			log.Fatal(err)
		}

		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &title, &content)
			if err != nil {
				log.Fatal(err)
			}

			res = append(res, dto.Article{ Id: id, Title: title, Content: content })
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()
		util.RespondTemplate(response, http.StatusOK, "template/index.html", res)

	// Default to 404
	} else {
		util.RespondNotFound(response)
	}
}

func PageView(response http.ResponseWriter, request *http.Request) {
	// Article page
	if(request.URL.Path[1:] == "view/1") {
		article1 := dto.Article{ Id: 1, Title: "First Article", Content: "Hello, World!" }
		util.RespondTemplate(response, http.StatusOK, "template/view.html", article1)
	
	// Default to 404
	} else {
		util.RespondNotFound(response)
	}
}


// --- Main ---

func main() {
	http.HandleFunc("/view/", PageView)
	http.HandleFunc("/", PageIndex)
	http.ListenAndServe(":8100", nil)
}