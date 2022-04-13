package tasks

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
	"github.com/GIT_USER_ID/GIT_REPO_ID/entities/task"
)

func UpdateTask(userId string, updatedTask task.TypeTask) task.TypeTask {
	var db = db.GetInstance()
	row := db.QueryRow("update tasks set (id=$1, title=$2, checked=$3, description=$4) where id = $1 and userId = $5 RETURNING *", updatedTask.Id, updatedTask.Title, updatedTask.Checked, updatedTask.Description, userId)
	var t task.TypeTask
	row.Scan(&t.UserId, &t.Id, &t.Checked, &t.Description, &t.Title)
	return task.TypeTask{Id: t.Id, UserId: t.UserId, Title: t.Title, Checked: t.Checked, Description: t.Description}
}
