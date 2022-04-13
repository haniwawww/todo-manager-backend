package tasks

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
	"github.com/GIT_USER_ID/GIT_REPO_ID/entities/task"
)

func GetTask(userId string, taskId string) task.TypeTask {
	var db = db.GetInstance()
	row := db.QueryRow("select * from tasks where userId = $1 and id = $2", userId, taskId)
	var t task.TypeTask
	// TODO 並び替えに対応できるようにする（せめてエラーが出るようにする）
	// マイグレーションした時に並びが変わると途端にバグる
	row.Scan(&t.Id, &t.Title, &t.Checked, &t.Description, &t.UserId)
	return task.TypeTask{Id: t.Id, Description: t.Description, Checked: t.Checked, Title: t.Title, UserId: t.UserId}
}
