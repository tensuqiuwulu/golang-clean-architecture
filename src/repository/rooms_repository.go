package repository

import (
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/entity"

	"gorm.io/gorm"
)

type RoomsRepository interface {
	GetRoomsById(db *gorm.DB, id string) (entity.RoomsModel, error)
}

type roomsRepositoryImpl struct{}

func NewRoomsRepository() RoomsRepository {
	return &roomsRepositoryImpl{}
}

func (r *roomsRepositoryImpl) GetRoomsById(db *gorm.DB, id string) (entity.RoomsModel, error) {
	var rooms entity.RoomsModel
	err := db.
		Table("rooms").
		Where("id = ?", id).First(&rooms).Error
	if err != nil {
		return rooms, err
	}
	return rooms, nil
}
