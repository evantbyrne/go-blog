package main

import (
	_ "github.com/lib/pq"
	"html/template"
	"io"
	"net/http"
    "database/sql"
    "log"
)

// --- Types ---

type Article struct {
	Id int
	Title string
	Content string
}


// --- Respond ---

func Respond(response http.ResponseWriter, status int, html string) {
	response.WriteHeader(status)
	io.WriteString(response, html)
	response.Header().Set("Content-Type", "text/html")
	response.Header().Set("Content-Length", string(len(html)))
}

func RespondNotFound(response http.ResponseWriter) {
	Respond(response, http.StatusNotFound, "<h1>Page Not Found</h1>")
}

func RespondTemplate(response http.ResponseWriter, status int, template_file string, data interface{}) {
	response.WriteHeader(status)
	t, _ := template.ParseFiles(template_file)
	t.Execute(response, data)
}


// --- Pages ---

func PageIndex(response http.ResponseWriter, request *http.Request) {
	// Index
	if (request.URL.Path == "/") {
		db, err := sql.Open("postgres", "user=evantbyrne dbname=go_blog sslmode=disable")
		if (err != nil) {
			log.Fatal(err)
		}

		res := make([]Article, 0)
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

			res = append(res, Article{ Id: id, Title: title, Content: content })
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()
		RespondTemplate(response, http.StatusOK, "template/index.html", res)

	// Default to 404
	} else {
		RespondNotFound(response)
	}
}

func PageView(response http.ResponseWriter, request *http.Request) {
	// Article page
	if(request.URL.Path[1:] == "view/1") {
		article1 := Article{ Id: 1, Title: "First Article", Content: "Hello, World!" }
		RespondTemplate(response, http.StatusOK, "template/view.html", article1)
	
	// Default to 404
	} else {
		RespondNotFound(response)
	}
}


// --- Main ---

func main() {
	http.HandleFunc("/view/", PageView)
	http.HandleFunc("/", PageIndex)
	http.ListenAndServe(":8100", nil)
}