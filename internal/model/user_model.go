package model

import "github.com/pocketbase/pocketbase/models"

type UserModel struct {
	models.BaseModel
}

func (*UserModel) TableName() string {
	return "users"
}
