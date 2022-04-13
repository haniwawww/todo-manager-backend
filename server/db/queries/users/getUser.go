package users

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
	"github.com/GIT_USER_ID/GIT_REPO_ID/entities/user"
)

func GetUser(userId string) user.TypeUser {
	var db = db.GetInstance()
	row := db.QueryRow("select * from users where id = $1", userId)
	var u user.TypeUser
	row.Scan(&u.Id, &u.Email_address, &u.Phone_number, &u.Name)
	return user.TypeUser{Id: u.Id, Name: u.Name, Email_address: u.Email_address, Phone_number: u.Phone_number}
}
