package model

import "github.com/pocketbase/pocketbase/models"

type Method string

const (
	WhatsApp Method = "whatsapp"
	Telegram Method = "web"
	Web      Method = "web"
	Manual   Method = "manual"
)

type ProviderModel struct {
	models.BaseModel
	Panel      string `db:"panel"`
	Url        string `db:"url"`
	Alias      string `db:"alias"`
	Method     Method `db:"method"`
	MethodData string `db:"method_data"`
}

func (*ProviderModel) TableName() string {
	return "providers"
}
