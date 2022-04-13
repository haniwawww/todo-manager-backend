package tasks

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
	"github.com/GIT_USER_ID/GIT_REPO_ID/entities/task"
)

func AddTask(task task.TypeTask) {
	var db = db.GetInstance()
	db.QueryRow("INSERT INTO tasks(id, userId, title, checked, description) VALUES($1, $2, $3, $4, $5)", task.Id, task.UserId, task.Title, task.Checked, task.Description)
}
