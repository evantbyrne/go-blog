package dao

import "log"
import "strconv"

import "../dto"
import "../../util"

func Query(sqlQuery string, args ...interface{}) []dto.Article {
	var db = util.DatabaseConnect()
	var stmt, err = db.Prepare(sqlQuery)
	if (err != nil) {
		log.Fatal(err)
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		log.Fatal(err)
	}

	var res = make([]dto.Article, 0)
	var id int
	var title string
	var content string
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &title, &content)
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

	return res
}

func GetArticleById(id string) (dto.Article, bool) {
	var idInt, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		return dto.Article{}, false
	}

	var res = Query("select * from article where id = $1", idInt)
	if len(res) < 1 {
		return dto.Article{}, false
	}

	return res[0], true
}

func GetAllArticles() []dto.Article {
	return Query("select * from article order by id desc")
}