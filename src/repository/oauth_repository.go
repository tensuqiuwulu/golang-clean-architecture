package repository

import (
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/entity"

	"gorm.io/gorm"
)

type OauthRepository interface {
	CheckOauthAccessTokenByApiKey(db *gorm.DB, apiKey string) (bool, error)
}

type oauthRepositoryImpl struct{}

func NewOauthRepository() OauthRepository {
	return &oauthRepositoryImpl{}
}

func (o *oauthRepositoryImpl) CheckOauthAccessTokenByApiKey(db *gorm.DB, apiKey string) (bool, error) {
	oauthAccessTokensModel := &entity.OauthAccessTokensModel{}

	err := db.
		Table("oauth_access_tokens").
		Where("id = ?", apiKey).First(&oauthAccessTokensModel).Error
	if err != nil {
		return false, err
	}

	if oauthAccessTokensModel.Id == "" {
		return false, nil
	}

	return true, nil
}
