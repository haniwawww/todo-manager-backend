package users

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
	"github.com/GIT_USER_ID/GIT_REPO_ID/entities/user"
)

func AddUser(user user.TypeUser) {
	var dbInstance = db.GetInstance()
	dbInstance.QueryRow("INSERT INTO users(id, name, email_address, phone_number) VALUES($1, $2, $3, $4) RETURNING (id, name, email_address, phone_number)", user.Id, user.Name, user.Email_address, user.Phone_number)
}
