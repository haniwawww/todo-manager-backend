package tasks

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
	"github.com/GIT_USER_ID/GIT_REPO_ID/entities/task"
)

func GetTask(userId string, taskId string) task.TypeTask {
	var db = db.GetInstance()
	row := db.QueryRow("select * from tasks where userId = $1 and id = $2", userId, taskId)
	var t task.TypeTask
	row.Scan(&t.UserId, &t.Id, &t.Checked, &t.Description, &t.Title)
	return task.TypeTask{Id: t.Id, UserId: t.UserId, Title: t.Title, Checked: t.Checked, Description: t.Description}
}
