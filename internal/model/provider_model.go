package model

import (
	"encoding/json"
	"fmt"

	"github.com/pocketbase/pocketbase/models"
)

type Method string

const (
	WhatsApp Method = "whatsapp"
	Telegram Method = "web"
	Web      Method = "web"
	Manual   Method = "manual"
)

type ProviderModel struct {
	models.BaseModel
	Panel              string `db:"panel"`
	Url                string `db:"url"`
	Alias              string `db:"alias"`
	Method             Method `db:"method"`
	MethodData         string `db:"method_data"`
	MethodDataReadable string `json:"-"`
}

func (*ProviderModel) TableName() string {
	return "providers"
}

func (p *ProviderModel) methodDataToJSON() map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(p.MethodData), &result)

	if err != nil {
		fmt.Println("err")
	}

	return result
}

func (p *ProviderModel) SetMethodDataReadableFromJsonString(methodData map[string]interface{}) {
	if p.Method == "telegram" {
		p.MethodDataReadable = methodData["telegram_chat_id"].(string)
	} else if p.Method == "web" {
		p.MethodDataReadable = methodData["support_username"].(string) + ":" + methodData["support_password"].(string)
	}
}

func (p *ProviderModel) SetMethodDataReadableFromSelf() {
	methodData := p.methodDataToJSON()

	if p.Method == "telegram" {
		p.MethodDataReadable = methodData["telegram_chat_id"].(string)
	} else if p.Method == "web" {
		p.MethodDataReadable = methodData["support_username"].(string) + ":" + methodData["support_password"].(string)
	}
}
