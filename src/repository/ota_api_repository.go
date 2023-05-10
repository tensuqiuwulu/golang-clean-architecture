package repository

import (
	"errors"

	"github.com/tensuqiuwulu/golang-clean-architecture/exception"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/entity"
	"github.com/tensuqiuwulu/golang-clean-architecture/utils"

	"github.com/labstack/echo/v4"
)

type OtaApiRepository interface {
	FetchRoomsAndRatesFromOta(c echo.Context, otaHotelId, otaHotelPassword, otaName string) []entity.OtaRemoteRoomRatesModel
	PushRoomsAndRatesToOta(c echo.Context, otaHotelId, otaHotelPassword, otaName string, otaRemoteRoomRatesModel []entity.OtaRemoteRoomRatesModel) error
}

type otaApiRepositoryImpl struct{}

func NewOtaApiRepository() OtaApiRepository {
	return &otaApiRepositoryImpl{}
}

func (o *otaApiRepositoryImpl) FetchRoomsAndRatesFromOta(c echo.Context, otaHotelId, otaHotelPassword, otaName string) []entity.OtaRemoteRoomRatesModel {
	switch otaName {
	case "Traveloka":
		otaRemoteRoomRatesModel := []entity.OtaRemoteRoomRatesModel{
			{
				RoomCode:     "112233",
				RoomRateCode: "223344",
				RoomType:     "Deluxe",
				RoomRateName: "Deluxe Rate",
			},
			{
				RoomCode:     "445566",
				RoomRateCode: "556677",
				RoomType:     "Superior",
				RoomRateName: "Superior Rate",
			},
			{
				RoomCode:     "778899",
				RoomRateCode: "223344",
				RoomType:     "Superior King",
				RoomRateName: "Deluxe Rate",
			},
			{
				RoomCode:     "225566",
				RoomRateCode: "556677",
				RoomType:     "Superior",
				RoomRateName: "Superior Rate",
			},
		}
		return otaRemoteRoomRatesModel

	case "Booking":
		otaRemoteRoomRatesModel := []entity.OtaRemoteRoomRatesModel{
			{
				RoomCode:     "112233",
				RoomRateCode: "223344",
				RoomType:     "Deluxe",
				RoomRateName: "Deluxe Rate",
			},
			{
				RoomCode:     "445566",
				RoomRateCode: "556677",
				RoomType:     "Superior",
				RoomRateName: "Superior Rate",
			},
			{
				RoomCode:     "778899",
				RoomRateCode: "223344",
				RoomType:     "Superior King",
				RoomRateName: "Deluxe Rate",
			},
			{
				RoomCode:     "225566",
				RoomRateCode: "556677",
				RoomType:     "Superior",
				RoomRateName: "Superior Rate",
			},
		}
		return otaRemoteRoomRatesModel

	case "Tiket":
		otaRemoteRoomRatesModel := []entity.OtaRemoteRoomRatesModel{
			{
				RoomCode:     "112233",
				RoomRateCode: "223344",
				RoomType:     "Deluxe",
				RoomRateName: "Deluxe Rate",
			},
			{
				RoomCode:     "445566",
				RoomRateCode: "556677",
				RoomType:     "Superior",
				RoomRateName: "Superior Rate",
			},
			{
				RoomCode:     "778899",
				RoomRateCode: "223344",
				RoomType:     "Superior King",
				RoomRateName: "Deluxe Rate",
			},
			{
				RoomCode:     "225566",
				RoomRateCode: "556677",
				RoomType:     "Superior",
				RoomRateName: "Superior Rate",
			},
		}
		return otaRemoteRoomRatesModel

	}
	exception.ErrorHandler(utils.LogStruct{Code: 404, Mssg: "OTA " + otaName + " Not available for connection", LogDetail: "OTA " + otaName + " Not available for connection", Data: nil})
	return nil
}

func (o *otaApiRepositoryImpl) PushRoomsAndRatesToOta(c echo.Context, otaHotelId, otaHotelPassword, otaName string, otaRemoteRoomRatesModel []entity.OtaRemoteRoomRatesModel) error {
	switch otaName {
	case "Traveloka":
		return nil
	}

	return errors.New("ota host not found")
}
