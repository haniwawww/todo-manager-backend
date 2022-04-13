package task

import "github.com/google/uuid"

type TypeTask struct {
	Id          string
	UserId      string
	Title       string
	Checked     bool
	Description string
}

type TypeReqTask struct {
	Title       string
	Checked     bool
	Description string
}

var instance *TypeTask

func GetInstance(userId string, self TypeReqTask) TypeTask {
	id := uuid.New()
	return TypeTask{
		Id:          id.String(),
		UserId:      userId,
		Title:       self.Title,
		Checked:     self.Checked,
		Description: self.Description,
	}
}
