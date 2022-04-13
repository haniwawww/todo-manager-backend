package users

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
)

func DeleteUser(userId string) {
	var db = db.GetInstance()
	db.QueryRow("delete from users where id = $1", userId)
}
