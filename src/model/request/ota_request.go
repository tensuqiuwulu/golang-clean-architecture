package request

type PullRoomsAndRatesFromOtaRequest struct {
	OtaPropertiesId        string `json:"ota_properties_id" form:"ota_properties_id" validate:"required"`
	OtaPropertiesPassword  string `json:"ota_properties_password" form:"ota_properties_password" validate:"required"`
	OtaPropertiesMappingId string `json:"ota_properties_mapping_id" form:"ota_properties_mapping_id" validate:"required"`
}

type PushRoomsAndRatesToOtaRequest struct {
	PropertiesId string `json:"property_id" form:"property_id" validate:"required"`
	PushType     string `json:"push_type" form:"push_type" validate:"required"`
}

type UpdateRoomsAvailabilitiesRequest struct {
	CustomerName           string  `json:"customer_name" form:"customer_name" validate:"required"`
	CustomerAddress        string  `json:"customer_address" form:"customer_address" validate:"required"`
	CustomerPhone          string  `json:"customer_phone" form:"customer_phone" validate:"required"`
	CustomerEmail          string  `json:"customer_email" form:"customer_email" validate:"required"`
	CustomerIdentityNumber string  `json:"customer_identity_number" form:"customer_identity_number" validate:"required"`
	RoomCode               string  `json:"room_code" form:"room_code" validate:"required"`
	FromDate               string  `json:"from_date" form:"from_date" validate:"required"`
	ToDate                 string  `json:"to_date" form:"to_date" validate:"required"`
	TotalNights            int     `json:"total_nights" form:"total_nights" validate:"required"`
	NumOfAdults            int     `json:"num_of_adults" form:"num_of_adults"`
	NumOfChildrens         int     `json:"num_of_childrens" form:"num_of_childrens"`
	NumOfRooms             int     `json:"num_of_rooms" form:"num_of_rooms" validate:"required"`
	TotalPrice             float64 `json:"total_price" form:"total_price" validate:"required"`
	OtaId                  string  `json:"ota_id" form:"ota_id" validate:"required"`
}
