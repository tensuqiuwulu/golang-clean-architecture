package response

import "github.com/tensuqiuwulu/golang-clean-architecture/src/model/entity"

type PullRoomsAndRatesFromOtaResponse struct {
	OtaName     string                           `json:"ota_name"`
	RemoteRooms []entity.OtaRemoteRoomRatesModel `json:"remote_rooms"`
}

type PushRatesAndAvailabilitiesToOtaResponse struct {
	PushType   string `json:"push_type"`
	StatusCode int    `json:"status_code"`
}

type TestRoomsAvailabilitesResponse struct {
	RoomsId    string `json:"rooms_id"`
	Date       string `json:"date"`
	NumOfRooms int    `json:"num_of_rooms"`
}
