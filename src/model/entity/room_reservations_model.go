package entity

import "time"

type RoomReservationsModel struct {
	Id          string    `gorm:"primaryKey;column:id;"`
	BookingsId  string    `gorm:"column:bookings_id;"`
	RoomsId     string    `gorm:"column:rooms_id;"`
	RatePlansId string    `gorm:"column:rate_plans_id;"`
	RoomPrice   float64   `gorm:"column:room_price;"`
	NumOfRooms  int       `gorm:"column:number_of_rooms;"`
	TotalPrice  float64   `gorm:"column:total_price;"`
	CreatedAt   time.Time `gorm:"column:created_at;"`
}
