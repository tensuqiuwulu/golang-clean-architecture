package repository

import (
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/entity"

	"gorm.io/gorm"
)

type MappingRepository interface {
	GetOtaRoomMappingByOtaRoomCode(db *gorm.DB, otaRoomCode string) (entity.OtaRoomMappingModel, error)
	GetOtaPropertiesMappingById(db *gorm.DB, otaPropertiesMappingId string) (entity.OtaPropertiesMappingModel, error)
	GetOtaRoomMappingByOtaPropertiesMappingId(db *gorm.DB, otaPropertiesMappingId string) ([]entity.OtaRoomMappingModel, error)
	CreateOtaRoomMapping(db *gorm.DB, otaRoomMapping []entity.OtaRoomMappingModel) error
}

type mappingRepositoryImpl struct{}

func NewMappingRepository() MappingRepository {
	return &mappingRepositoryImpl{}
}

func (m *mappingRepositoryImpl) GetOtaRoomMappingByOtaRoomCode(db *gorm.DB, otaRoomCode string) (entity.OtaRoomMappingModel, error) {
	var otaRoomMapping entity.OtaRoomMappingModel
	err := db.
		Table("ota_room_mapping").
		Where("ota_rooms_code = ?", otaRoomCode).Find(&otaRoomMapping).Error
	if err != nil {
		return otaRoomMapping, err
	}
	return otaRoomMapping, nil
}

func (m *mappingRepositoryImpl) GetOtaPropertiesMappingById(db *gorm.DB, otaPropertiesMappingId string) (entity.OtaPropertiesMappingModel, error) {
	var otaPropertiesMapping entity.OtaPropertiesMappingModel
	err := db.
		Table("ota_properties_mapping").
		Where("id = ?", otaPropertiesMappingId).Find(&otaPropertiesMapping).Error
	if err != nil {
		return otaPropertiesMapping, err
	}
	return otaPropertiesMapping, nil
}

func (m *mappingRepositoryImpl) GetOtaRoomMappingByOtaPropertiesMappingId(db *gorm.DB, otaPropertiesMappingId string) ([]entity.OtaRoomMappingModel, error) {
	var otaRoomMapping []entity.OtaRoomMappingModel
	err := db.
		Table("ota_room_mapping").
		Where("ota_properties_mapping_id = ?", otaPropertiesMappingId).Find(&otaRoomMapping).Error
	if err != nil {
		return otaRoomMapping, err
	}
	return otaRoomMapping, nil
}

func (m *mappingRepositoryImpl) CreateOtaRoomMapping(db *gorm.DB, otaRoomMapping []entity.OtaRoomMappingModel) error {
	err := db.Table("ota_room_mapping").Create(&otaRoomMapping).Error
	if err != nil {
		return err
	}
	return nil
}
