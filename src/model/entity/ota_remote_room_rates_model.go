package entity

type OtaRemoteRoomRatesModel struct {
	RoomCode     string `json:"room_code"`
	RoomRateCode string `json:"room_rate_code"`
	RoomType     string `json:"room_type"`
	RoomRateName string `json:"room_rate_name"`
}
