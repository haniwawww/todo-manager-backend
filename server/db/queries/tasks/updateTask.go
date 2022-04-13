package tasks

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
	"github.com/GIT_USER_ID/GIT_REPO_ID/entities/task"
)

func UpdateTask(userId string, task task.TypeTask) {
	var db = db.GetInstance()
	db.QueryRow("update tasks set (id=$1, title=$2, checked=$3, description=$4) where id = $1 and userId = $5", task.Id, task.Title, task.Checked, task.Description, userId)
}
