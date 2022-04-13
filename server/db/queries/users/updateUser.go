package users

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
	"github.com/GIT_USER_ID/GIT_REPO_ID/entities/user"
)

func UpdateUser(userId string, user user.TypeUser) {
	var db = db.GetInstance()
	db.QueryRow("update users set (name = $1, email_address = $2, phone_number=$3) where id = $4", user.Name, user.Email_address, user.Phone_number, userId)
}
