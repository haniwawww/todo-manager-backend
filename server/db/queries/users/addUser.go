package users

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
	"github.com/GIT_USER_ID/GIT_REPO_ID/entities/user"
)

func AddUser(newUser user.TypeUser) user.TypeUser {
	var dbInstance = db.GetInstance()
	row := dbInstance.QueryRow("INSERT INTO users(id, name, email_address, phone_number) VALUES($1, $2, $3, $4) RETURNING *", newUser.Id, newUser.Name, newUser.Email_address, newUser.Phone_number)
	var u user.TypeUser
	row.Scan(&u.Id, &u.Email_address, &u.Phone_number, &u.Name)
	return user.TypeUser{Id: u.Id, Name: u.Name, Email_address: u.Email_address, Phone_number: u.Phone_number}
}
