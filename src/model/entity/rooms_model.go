package entity

type RoomsModel struct {
	Id               string  `gorm:"primaryKey;column:id;"`
	PropertiesId     string  `gorm:"column:properties_id;"`
	RoomsTypesId     string  `gorm:"column:rooms_types_id;"`
	MaxOccupancy     int     `gorm:"column:max_occupancy;"`
	BedTypesId       string  `gorm:"column:bed_types_id;"`
	MaxExtraBeds     int     `gorm:"column:max_extra_beds;"`
	PriceExtraBeds   float64 `gorm:"column:price_extra_beds;"`
	RoomSize         int     `gorm:"column:room_size;"`
	Breakfast        int     `gorm:"column:breakfast;"`
	RoomRatePriceOta float64 `gorm:"column:room_rate_price_ota;"`
	ExtraGuestFee    float64 `gorm:"column:extra_guest_fee;"`
	NumOfRooms       int     `gorm:"column:number_of_rooms;"`
	RoomName         string  `gorm:"column:room_name;"`
}
