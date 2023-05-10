package entity

import "time"

type CustomerModel struct {
	Id             string    `gorm:"primaryKey;column:id;"`
	UsersId        string    `gorm:"column:users_id;"`
	Name           string    `gorm:"column:name;"`
	Address        string    `gorm:"column:address;"`
	PhoneNumber    string    `gorm:"column:phone_number;"`
	Email          string    `gorm:"column:email;"`
	Gender         int       `gorm:"column:gender;"`
	Age            int       `gorm:"column:age;"`
	Status         int       `gorm:"column:status;"`
	IdentityNumber string    `gorm:"column:identity_number;"`
	CreatedAt      time.Time `gorm:"column:created_at;"`
}
