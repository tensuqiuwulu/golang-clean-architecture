package entity

import "gorm.io/gorm"

type RatesAvailabilitiesModel struct {
	gorm.Model
	Id               string  `gorm:"primaryKey;column:id;"`
	PropertiesId     string  `gorm:"column:properties_id;"`
	RoomsId          string  `gorm:"column:rooms_id;"`
	Date             string  `gorm:"column:date;"`
	NumberOfRooms    int     `gorm:"column:number_of_rooms;"`
	MaxOccupancy     int     `gorm:"column:max_occupancy;"`
	RoomType         string  `gorm:"column:room_type;"`
	PropertyName     string  `gorm:"column:property_name;"`
	Country          string  `gorm:"column:country;"`
	Province         string  `gorm:"column:province;"`
	IsAllRooms       int     `gorm:"column:is_all_rooms;"`
	OtaId            string  `gorm:"column:ota_id;"`
	OtaNumberOfRooms int     `gorm:"column:ota_number_of_rooms;"`
	RatePlansId      string  `gorm:"column:rate_plans_id;"`
	PlanName         string  `gorm:"column:plan_name;"`
	OtaRate          float64 `gorm:"column:ota_rate;"`
	Coa              int     `gorm:"column:coa;"`
	Cod              int     `gorm:"column:cod;"`
	StopSell         int     `gorm:"column:stop_sell;"`
	StopSellLimit    int     `gorm:"column:stop_sell_limit;"`
	CreatedAt        string  `gorm:"column:created_at;"`
	UpdatedAt        string  `gorm:"column:updated_at;"`
	IsPush           int     `gorm:"column:is_push"`
	TableName        string  `gorm:"-" json:"-"` // Ignore this field
	Version          int
}
