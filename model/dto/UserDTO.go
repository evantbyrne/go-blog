package dto

type User struct {
	Email string `db:"email"`
	Password string `db:"password"`
}