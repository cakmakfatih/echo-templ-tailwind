package entity

import "time"

type PanelSession struct {
	Id              string
	User            string
	LoginURL        string
	SupportUsername string
	SupportPassword string
	TelegramToken   string
	WhatsappToken   string
	Created         time.Time
	Updated         time.Time
}
