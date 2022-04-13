package users

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
	"github.com/GIT_USER_ID/GIT_REPO_ID/entities/user"
)

func UpdateUser(userId string, updatedUser user.TypeUser) user.TypeUser {
	var db = db.GetInstance()
	row := db.QueryRow("update users set (name = $1, email_address = $2, phone_number=$3) where id = $4 RETURNING *", updatedUser.Name, updatedUser.Email_address, updatedUser.Phone_number, userId)
	var u user.TypeUser
	row.Scan(&u.Id, &u.Email_address, &u.Phone_number, &u.Name)
	return user.TypeUser{Id: u.Id, Name: u.Name, Email_address: u.Email_address, Phone_number: u.Phone_number}
}
