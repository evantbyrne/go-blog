package dao

import "../dto"
import "../../util"

func CheckUser(email string, hashed_password string) bool {
	var db = new(util.Database).Gorp()
	var res dto.User
	err := db.SelectOne(&res, "select * from \"user\" where email = $1 and password = $2", email, hashed_password)
	return err == nil
}