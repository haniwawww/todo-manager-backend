package user

import "github.com/google/uuid"

type TypeUser struct {
	Id            string
	Name          string
	Email_address string
	Phone_number  string
}

type TypeReqUser struct {
	Name         string
	EmailAddress string
	PhoneNumber  string
}

func GetInstance(self TypeReqUser) TypeUser {
	id := uuid.New()
	return TypeUser{
		Id:            id.String(),
		Name:          self.Name,
		Email_address: self.EmailAddress,
		Phone_number:  self.PhoneNumber,
	}
}
