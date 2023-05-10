package repository

import (
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/entity"

	"gorm.io/gorm"
)

type OtaRepository interface {
	GetOtaById(db *gorm.DB, id string) (entity.OtaModel, error)
}

type otaRepositoryImpl struct{}

func NewOtaRepository() OtaRepository {
	return &otaRepositoryImpl{}
}

func (o *otaRepositoryImpl) GetOtaById(db *gorm.DB, id string) (entity.OtaModel, error) {
	var ota entity.OtaModel
	err := db.
		Table("ota").
		Where("id = ?", id).First(&ota).Error
	if err != nil {
		return ota, err
	}
	return ota, nil
}
