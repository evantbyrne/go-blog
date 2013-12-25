package main

import "net/http"

import "github.com/gorilla/mux"

import "./app/model/dao"
import "./app/model/dto"
import "./app/util"


// --- Pages ---

func PageNotFound(response http.ResponseWriter, request *http.Request) {
	util.RespondNotFound(response)
}

func PageIndex(response http.ResponseWriter, request *http.Request) {
	res := dao.GetAllArticles()
	util.RespondTemplate(response, http.StatusOK, "template/index.html", res)
}

func PageView(response http.ResponseWriter, request *http.Request) {
	var vars = mux.Vars(request)
	var id = vars["id"]
	var res, success = dao.GetArticleById(id)
	if success {
		util.RespondTemplate(response, http.StatusOK, "template/view.html", res)
	} else {
		util.RespondNotFound(response)
	}
}

func PageEdit(response http.ResponseWriter, request *http.Request) {
	var vars = mux.Vars(request)
	var id = vars["id"]
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
	var vars = mux.Vars(request)
	var id = vars["id"]
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
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(PageNotFound)
	r.HandleFunc("/", PageIndex)
	r.HandleFunc("/view/{id:[0-9]+}", PageView)
	r.HandleFunc("/edit/{id:[0-9]+}", PageEdit)
	r.HandleFunc("/delete/{id:[0-9]+}", PageDelete)
	r.HandleFunc("/create", PageCreate)
	http.Handle("/", r)
	http.ListenAndServe(":8100", nil)
}