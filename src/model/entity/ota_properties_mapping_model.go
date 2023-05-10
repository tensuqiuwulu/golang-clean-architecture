package entity

type OtaPropertiesMappingModel struct {
	Id                    string `gorm:"primaryKey;column:id;"`
	OtaId                 string `gorm:"column:ota_id;"`
	PropertiesId          string `gorm:"column:properties_id;"`
	OtaPropertiesCode     string `gorm:"column:ota_properties_code;"`
	OtaPropertiesPassword string `gorm:"column:ota_properties_password;"`
}
