package main

import "net/http"

import "./app/model/dao"
import "./app/util"

// --- Pages ---

func PageIndex(response http.ResponseWriter, request *http.Request) {
	// Index
	if (request.URL.Path == "/") {
		res := dao.GetAllArticles()
		util.RespondTemplate(response, http.StatusOK, "template/index.html", res)
	} else {
		util.RespondNotFound(response)
	}
}

func PageView(response http.ResponseWriter, request *http.Request) {
	// Article page
	var id = request.URL.Path[6:]
	var res, success = dao.GetArticleById(id)
	if success {
		util.RespondTemplate(response, http.StatusOK, "template/view.html", res)
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