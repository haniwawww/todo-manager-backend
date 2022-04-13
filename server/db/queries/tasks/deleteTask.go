package tasks

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
)

func DeleteTask(userId string, taskId string) {
	var db = db.GetInstance()
	db.QueryRow("delete from tasks where userId = $1 and taskId = $2", userId, taskId)
}
