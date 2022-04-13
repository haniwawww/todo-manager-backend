package tasks

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/db"
	"github.com/GIT_USER_ID/GIT_REPO_ID/entities/task"
)

func GetTasks(userId string) ([]task.TypeTask, error) {
	var db = db.GetInstance()
	rows, err := db.Query("select * from tasks where userId = $1", userId)
	var ts []task.TypeTask
	for rows.Next() {
		var t task.TypeTask
		// TODO 並び替えに対応できるようにする（せめてエラーが出るようにする）
		// マイグレーションした時に並びが変わると途端にバグる
		rows.Scan(&t.Id, &t.Title, &t.Checked, &t.Description, &t.UserId)
		ts = append(ts, t)
	}
	return ts, err
}
