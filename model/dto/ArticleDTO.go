package dto

type Article struct {
	Id int64 `db:"id"`
	Title string `db:"title"`
	Content string `db:"content"`
}