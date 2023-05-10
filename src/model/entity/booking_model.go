package entity

import "time"

type BookingModel struct {
	Id                string    `gorm:"primaryKey;column:id;"`
	PropertiesId      string    `gorm:"column:properties_id;"`
	CustomersId       string    `gorm:"column:customers_id;"`
	MicesId           string    `gorm:"column:mices_id; default:null;"`
	BookingCode       string    `gorm:"column:booking_code;"`
	BookingDate       time.Time `gorm:"column:booking_date;"`
	StartDate         time.Time `gorm:"column:start_date;"`
	EndDate           time.Time `gorm:"column:end_date;"`
	BookingType       string    `gorm:"column:booking_type;"`
	AgentsId          string    `gorm:"column:agents_id; default:null;"`
	OtaId             string    `gorm:"column:ota_id;"`
	NumberOfAdults    int       `gorm:"column:number_of_adults;"`
	NumberOfChildrens int       `gorm:"column:number_of_childrens;"`
	ExtraBeds         int       `gorm:"column:extra_beds;"`
	TotalPrice        float64   `gorm:"column:total_price;"`
	Note              string    `gorm:"column:note;"`
	Status            int       `gorm:"column:status;"`
	CreatedAt         time.Time `gorm:"column:created_at;"`
}
