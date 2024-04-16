package model

import "github.com/pocketbase/pocketbase/models"

type ServiceModel struct {
	models.BaseModel
	Provider       string `db:"provider"`
	ServiceId      int    `db:"service_id"`
	Name           string `db:"name"`
	HasRefill      bool   `db:"has_refill"`
	RefillDuration int    `db:"refill_duration"`
}

func (*ServiceModel) TableName() string {
	return "services"
}
