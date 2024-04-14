package model

import "github.com/pocketbase/pocketbase/models"

type UserModel struct {
	models.BaseModel
	Email    string
	Username string
	Roles    []string
}

func (*UserModel) TableName() string {
	return "users"
}
