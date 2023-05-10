package repository

import (
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/entity"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(db *gorm.DB, booking *entity.BookingModel, tableName string) error
}

type bookingRepositoryImpl struct{}

func NewBookingRepository() BookingRepository {
	return &bookingRepositoryImpl{}
}

func (r *bookingRepositoryImpl) CreateBooking(db *gorm.DB, booking *entity.BookingModel, tableName string) error {
	err := db.Table(tableName).Create(booking).Error
	if err != nil {
		return err
	}
	return nil
}
