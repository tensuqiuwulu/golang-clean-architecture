package repository

import (
	"errors"
	"time"

	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RatesAvailabilitiesRepository interface {
	GetRatesAvailabilitesByDate(c echo.Context, db *gorm.DB, numOfRooms int, roomId, dateIn, dateOut, tableName string) ([]entity.RatesAvailabilitiesModel, error)
	GetRatesAvailabilitiesNotPushed(db *gorm.DB, propertiesId, tableName string) ([]entity.RatesAvailabilitiesModel, error)
	UpdateStatusRatesAvailabilitiesPushed(db *gorm.DB, ratesAvailabilitiesId []string, tableName string) error
	UpdateAvailabilities(db *gorm.DB, numOfRooms int, ratesAvailabilitiesId []string, tableName string) error
}

type ratesAvailabilitiesRepositoryImpl struct{}

func NewRatesAvailabilitiesRepository() RatesAvailabilitiesRepository {
	return &ratesAvailabilitiesRepositoryImpl{}
}

func (r *ratesAvailabilitiesRepositoryImpl) GetRatesAvailabilitesByDate(c echo.Context, db *gorm.DB, numOfRooms int, roomId, dateIn, dateOut, tableName string) ([]entity.RatesAvailabilitiesModel, error) {

	var ratesAvailabilities []entity.RatesAvailabilitiesModel
	err := db.
		Table(tableName).
		Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		}).
		Set("gorm:lock_timeout", 10*time.Second).
		Where("rooms_id = ? AND number_of_rooms >= ?", roomId, dateIn, dateOut, numOfRooms).
		Where("date => ?", dateIn).
		Where("date < ?", dateOut).
		Order("date ASC").
		Find(&ratesAvailabilities).Error

	if err != nil {
		return ratesAvailabilities, err
	}

	for i := range ratesAvailabilities {
		ratesAvailabilities[i].TableName = tableName
	}

	return ratesAvailabilities, nil

}

func (r *ratesAvailabilitiesRepositoryImpl) GetRatesAvailabilitiesNotPushed(db *gorm.DB, propertiesId, tableName string) ([]entity.RatesAvailabilitiesModel, error) {
	var ratesAvailabilities []entity.RatesAvailabilitiesModel
	err := db.
		Table(tableName).
		Where("properties_id = ? AND is_push = ?", propertiesId, 0).Find(&ratesAvailabilities).Error
	if err != nil {
		return ratesAvailabilities, err
	}

	for i := range ratesAvailabilities {
		ratesAvailabilities[i].TableName = tableName
	}

	return ratesAvailabilities, nil
}

func (r *ratesAvailabilitiesRepositoryImpl) UpdateStatusRatesAvailabilitiesPushed(db *gorm.DB, ratesAvailabilitiesId []string, tableName string) error {
	err := db.Table(tableName).Where("id IN ?", ratesAvailabilitiesId).Update("is_push", 1).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ratesAvailabilitiesRepositoryImpl) UpdateAvailabilities(db *gorm.DB, numOfRooms int, ratesAvailabilitiesId []string, tableName string) error {
	err := db.
		Table(tableName).
		Where("id IN ?", ratesAvailabilitiesId).
		Update("number_of_rooms", gorm.Expr("number_of_rooms - ?", numOfRooms)).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ratesAvailabilitiesRepositoryImpl) UpdateAvailabilitiesNonBatch(db *gorm.DB, numOfRooms int, ratesAvailabilitiesId, tableName string) error {
	ratesAvailabilities := entity.RatesAvailabilitiesModel{}

	err := db.
		Table(tableName).
		Model(&ratesAvailabilities). // <- this is important
		Where("id = ?", ratesAvailabilitiesId).
		Update("number_of_rooms", gorm.Expr("number_of_rooms - ?", numOfRooms)).Error
	if err != nil {
		return err
	}

	if ratesAvailabilities.NumberOfRooms < 0 {
		return errors.New("number of rooms is less than 0")
	}

	return nil
}
