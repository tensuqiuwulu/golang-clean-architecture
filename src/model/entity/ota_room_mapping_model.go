package entity

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type OtaRoomMappingModel struct {
	Id                     string    `gorm:"primaryKey;column:id;"`
	OtaPropertiesMappingId string    `gorm:"column:ota_properties_mapping_id;"`
	OtaRoomsCode           string    `gorm:"column:ota_rooms_code;"`
	OtaRatePlansCode       string    `gorm:"column:ota_rate_plans_code;"`
	OtaRatePlanName        string    `gorm:"column:ota_rate_plan_name;"`
	OtaRoomName            string    `gorm:"column:ota_room_name;"`
	RoomsId                string    `gorm:"column:rooms_id; default:null;"`
	RatePlansId            string    `gorm:"column:rate_plans_id; default:null;"`
	PlanName               string    `gorm:"column:plan_name;"`
	Status                 int       `gorm:"column:status;"`
	CreatedAt              time.Time `gorm:"column:created_at;"`
	UpdatedAt              null.Time `gorm:"column:updated_at;"`
}
