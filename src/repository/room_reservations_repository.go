package repository

import (
	"log"

	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/entity"

	"gorm.io/gorm"
)

type RoomReservationsRepository interface {
	CreateRoomReservations(db *gorm.DB, roomReservations *entity.RoomReservationsModel, tableName string) error
}

type roomReservationsRepositoryImpl struct{}

func NewRoomReservationsRepository() RoomReservationsRepository {
	return &roomReservationsRepositoryImpl{}
}

func (r *roomReservationsRepositoryImpl) CreateRoomReservations(db *gorm.DB, roomReservations *entity.RoomReservationsModel, tableName string) error {
	err := db.Table(tableName).Create(roomReservations).Error
	if err != nil {
		log.Println("error create room reservations", err.Error())
		return err
	}

	return nil
}
