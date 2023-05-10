package entity

type OtaModel struct {
	Id      string `gorm:"primaryKey;column:id;"`
	UsersId string `gorm:"column:users_id;"`
	Name    string `gorm:"column:name;"`
	ApiKey  string `gorm:"column:api_key;"`
}
