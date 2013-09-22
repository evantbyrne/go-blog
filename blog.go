package main

import (
	"html/template"
	"io"
	"net/http"
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
		article1 := Article{ Id: 1, Title: "First Article", Content: "Hello, World!" }
		RespondTemplate(response, http.StatusOK, "template/index.html", []*Article{ &article1 })

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