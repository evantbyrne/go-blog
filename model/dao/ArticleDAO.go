package dao

import "log"
import "strconv"

import "../dto"
import "../../util"

func GetArticleById(id string) (dto.Article, bool) {
	var idInt, err = strconv.ParseInt(id, 10, 64)
	if err == nil {
		var db = new(util.Database).Gorp()
		var res dto.Article;
		err = db.SelectOne(&res, "select * from article where id = $1", idInt)
		if err == nil {
			return res, true
		}
	}

	return dto.Article{}, false
}

func GetAllArticles() []dto.Article {
	var db = new(util.Database).Gorp()
	var res []dto.Article;
	_, err := db.Select(&res, "select * from article order by id desc")
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func UpdateArticle(article dto.Article) {
	var db = new(util.Database).Open()
	db.Exec("update article set title = $2, content = $3 where id = $1", article.Id, article.Title, article.Content)
}

func CreateArticle(article dto.Article) {
	var db = new(util.Database).Open()
	db.Exec("insert into article (title, content) values ($1, $2)", article.Title, article.Content)
}

func DeleteArticle(id int64) {
	var db = new(util.Database).Open()
	db.Exec("delete from article where id = $1", id)
}