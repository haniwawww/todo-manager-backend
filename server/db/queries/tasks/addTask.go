package tasks

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
	"github.com/GIT_USER_ID/GIT_REPO_ID/entities/task"
)

func AddTask(newTask task.TypeTask) task.TypeTask {
	var db = db.GetInstance()
	row := db.QueryRow("INSERT INTO tasks(id, userId, title, checked, description) VALUES($1, $2, $3, $4, $5) RETURNING id, userId, title, checked, description", newTask.Id, newTask.UserId, newTask.Title, newTask.Checked, newTask.Description)
	var t task.TypeTask
	row.Scan(&t.Id, &t.UserId, &t.Title, &t.Checked, &t.Description)
	return task.TypeTask{Id: t.Id, UserId: t.UserId, Title: t.Title, Checked: t.Checked, Description: t.Description}
}
