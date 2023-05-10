package entity

import "time"

type OauthAccessTokensModel struct {
	Id        string    `gorm:"primaryKey;column:id;"`
	UserId    string    `gorm:"column:user_id;"`
	ClientId  string    `gorm:"column:client_id;"`
	Name      string    `gorm:"column:name;"`
	Scopes    string    `gorm:"column:scopes;"`
	Revoked   int       `gorm:"column:revoked;"`
	ExpiresAt time.Time `gorm:"column:expires_at;"`
}
