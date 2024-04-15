package model

import (
	"time"

	"github.com/pocketbase/pocketbase/models"
)

type PanelModel struct {
	models.BaseModel
	User            string    `db:"user"`
	LoginURL        string    `db:"login_url"`
	SupportUsername string    `db:"support_username"`
	SupportPassword string    `db:"support_password"`
	TelegramToken   string    `db:"telegram_token"`
	WhatsappToken   string    `db:"whatsapp_token"`
	Created         time.Time `db:"created"`
	Updated         time.Time `db:"updated"`
}

func (*PanelModel) TableName() string {
	return "panels"
}
