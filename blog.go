package main

import "net/http"

import "./app/model/dao"
import "./app/model/dto"
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

func PageEdit(response http.ResponseWriter, request *http.Request) {
	var id = request.URL.Path[6:]
	var res, success = dao.GetArticleById(id)
	if success {
		if(request.Method == "POST") {
			request.ParseForm()
			res.Title = request.PostForm.Get("title")
			res.Content = request.PostForm.Get("content")
			dao.UpdateArticle(res)
			util.RespondTemplate(response, http.StatusOK, "template/view.html", res)
		} else {
			util.RespondTemplate(response, http.StatusOK, "template/edit.html", res)
		}
	} else {
		util.RespondNotFound(response)
	}
}

func PageDelete(response http.ResponseWriter, request *http.Request) {
	var id = request.URL.Path[8:]
	var res, success = dao.GetArticleById(id)
	if success {
		if(request.Method == "POST") {
			dao.DeleteArticle(res.Id)
			http.Redirect(response, request, "/", 303)
		} else {
			util.RespondTemplate(response, http.StatusOK, "template/delete.html", res)
		}
	} else {
		util.RespondNotFound(response)
	}
}

func PageCreate(response http.ResponseWriter, request *http.Request) {
	if(request.Method == "POST") {
		request.ParseForm()
		dao.CreateArticle(dto.Article{
			Title: request.PostForm.Get("title"),
			Content: request.PostForm.Get("content"),
		})
		http.Redirect(response, request, "/", 303)
	} else {
		util.RespondTemplate(response, http.StatusOK, "template/create.html", dto.Article{})
	}
}


// --- Main ---

func main() {
	http.HandleFunc("/view/", PageView)
	http.HandleFunc("/edit/", PageEdit)
	http.HandleFunc("/delete/", PageDelete)
	http.HandleFunc("/create/", PageCreate)
	http.HandleFunc("/", PageIndex)
	http.ListenAndServe(":8100", nil)
}