package service

import (
	"net/http"
	"sync"
	"time"

	"github.com/tensuqiuwulu/golang-clean-architecture/exception"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/entity"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/request"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/response"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/repository"
	"github.com/tensuqiuwulu/golang-clean-architecture/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OtaService interface {
	FetchRoomsAndRatesFromOta(c echo.Context, pullRoomsAndRatesFromOtaRequest *request.PullRoomsAndRatesFromOtaRequest) response.PullRoomsAndRatesFromOtaResponse
	PushRoomsAndRatesToOta(c echo.Context, pushRoomsAndRatesToOtaRequest *request.PushRoomsAndRatesToOtaRequest) response.PushRatesAndAvailabilitiesToOtaResponse
	UpdateRoomsAvailabilities(c echo.Context, updateRoomsAvailabilitiesRequest *request.UpdateRoomsAvailabilitiesRequest) interface{}
}

type otaServiceImpl struct {
	DB                            *gorm.DB
	MappingRepository             repository.MappingRepository
	OtaApiRepository              repository.OtaApiRepository
	OtaRepository                 repository.OtaRepository
	RatesAvailabilitiesRepository repository.RatesAvailabilitiesRepository
	CustomerRepository            repository.CustomerRepository
	BookingRepository             repository.BookingRepository
	RoomReservationsRepository    repository.RoomReservationsRepository
	RoomsRepository               repository.RoomsRepository
}

func NewOtaService(
	db *gorm.DB,
	mappingRepository repository.MappingRepository,
	otaApiRepository repository.OtaApiRepository,
	otaRepository repository.OtaRepository,
	ratesAvailabilitiesRepository repository.RatesAvailabilitiesRepository,
	customerRepository repository.CustomerRepository,
	bookingRepository repository.BookingRepository,
	roomReservationsRepository repository.RoomReservationsRepository,
	roomsRepository repository.RoomsRepository,
) OtaService {
	return &otaServiceImpl{
		DB:                            db,
		MappingRepository:             mappingRepository,
		OtaApiRepository:              otaApiRepository,
		OtaRepository:                 otaRepository,
		RatesAvailabilitiesRepository: ratesAvailabilitiesRepository,
		CustomerRepository:            customerRepository,
		BookingRepository:             bookingRepository,
		RoomReservationsRepository:    roomReservationsRepository,
		RoomsRepository:               roomsRepository,
	}
}

func (m *otaServiceImpl) FetchRoomsAndRatesFromOta(c echo.Context, pullRoomsAndRatesFromOtaRequest *request.PullRoomsAndRatesFromOtaRequest) response.PullRoomsAndRatesFromOtaResponse {
	var err error

	otaPropertiesMapping, err := m.MappingRepository.GetOtaPropertiesMappingById(m.DB, pullRoomsAndRatesFromOtaRequest.OtaPropertiesMappingId)
	if err != nil {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Get Ota Properties Mapping By Id " + err.Error()})
	}

	if len(otaPropertiesMapping.Id) == 0 {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusNotFound, Mssg: "Ota Properties Mapping Not Found", LogDetail: "Ota Properties Mapping Not Found"})
	}

	ota, err := m.OtaRepository.GetOtaById(m.DB, otaPropertiesMapping.OtaId)
	if err != nil {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Get Ota By Id " + err.Error()})
	}

	if len(ota.Id) == 0 {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusNotFound, Mssg: "Ota Not Found", LogDetail: "Ota Not Found"})
	}

	otaRemoteRoomsRatesModel := m.OtaApiRepository.FetchRoomsAndRatesFromOta(c, otaPropertiesMapping.OtaPropertiesCode, otaPropertiesMapping.OtaPropertiesPassword, ota.Name)

	otaRoomMapping, err := m.MappingRepository.GetOtaRoomMappingByOtaPropertiesMappingId(m.DB, otaPropertiesMapping.Id)
	if err != nil {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Get Ota Room Mapping By Ota Properties Mapping Id " + err.Error()})
	}

	response := response.PullRoomsAndRatesFromOtaResponse{
		OtaName:     ota.Name,
		RemoteRooms: otaRemoteRoomsRatesModel,
	}

	if len(otaRoomMapping) == 0 {
		otaRoomMappings := []entity.OtaRoomMappingModel{}
		for _, v := range otaRemoteRoomsRatesModel {
			otaRoomMappings = append(otaRoomMappings, entity.OtaRoomMappingModel{
				Id:                     utils.RandomUUID(),
				OtaPropertiesMappingId: otaPropertiesMapping.Id,
				OtaRoomsCode:           v.RoomCode,
				OtaRatePlansCode:       v.RoomRateCode,
				OtaRatePlanName:        v.RoomRateName,
				OtaRoomName:            v.RoomType,
				CreatedAt:              time.Now(),
			})
		}

		err = m.MappingRepository.CreateOtaRoomMapping(m.DB, otaRoomMappings)
		if err != nil {
			exception.ErrorHandler(utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Create Ota Room Mapping " + err.Error()})
		}

		return response
	} else {
		return response
	}
}

func (m *otaServiceImpl) PushRoomsAndRatesToOta(c echo.Context, pushRoomsAndRatesToOtaRequest *request.PushRoomsAndRatesToOtaRequest) response.PushRatesAndAvailabilitiesToOtaResponse {
	switch pushRoomsAndRatesToOtaRequest.PushType {
	case "availability":
		if mssg := m.pushAvailabilites(pushRoomsAndRatesToOtaRequest.PropertiesId); mssg != "" {
			return response.PushRatesAndAvailabilitiesToOtaResponse{
				PushType:   pushRoomsAndRatesToOtaRequest.PushType,
				StatusCode: 0,
			}
		}
	case "rate":
		if mssg := m.pushRates(pushRoomsAndRatesToOtaRequest.PropertiesId); mssg != "" {
			return response.PushRatesAndAvailabilitiesToOtaResponse{
				PushType:   pushRoomsAndRatesToOtaRequest.PushType,
				StatusCode: 0,
			}
		}
	case "all":
		if mssg := m.pushAll(pushRoomsAndRatesToOtaRequest.PropertiesId); mssg != "" {
			return response.PushRatesAndAvailabilitiesToOtaResponse{
				PushType:   pushRoomsAndRatesToOtaRequest.PushType,
				StatusCode: 0,
			}
		}
	default:
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusBadRequest, Mssg: "Push Type Not Found", LogDetail: "Push Type Not Found"})
	}

	return response.PushRatesAndAvailabilitiesToOtaResponse{
		PushType:   pushRoomsAndRatesToOtaRequest.PushType,
		StatusCode: 1,
	}
}

func (m *otaServiceImpl) pushAll(propertiesId string) string {
	roomRatesAvailabilities, mssg := m.getRatesAvailabilitesNotAlreadyPushed(propertiesId)
	if mssg != "" {
		return mssg
	}

	// push logic into ota

	// update status to pushed
	m.updateRatesAvailabilitesStatusPushed(roomRatesAvailabilities)

	return ""
}

func (m *otaServiceImpl) pushRates(propertiesId string) string {
	roomRatesAvailabilities, mssg := m.getRatesAvailabilitesNotAlreadyPushed(propertiesId)
	if mssg != "" {
		return mssg
	}

	// push logic into ota

	// update status to pushed
	m.updateRatesAvailabilitesStatusPushed(roomRatesAvailabilities)

	return ""
}

func (m *otaServiceImpl) pushAvailabilites(propertiesId string) string {
	roomRatesAvailabilities, mssg := m.getRatesAvailabilitesNotAlreadyPushed(propertiesId)
	if mssg != "" {
		return mssg
	}

	// push logic into ota

	// update status to pushed
	m.updateRatesAvailabilitesStatusPushed(roomRatesAvailabilities)

	return ""
}

func (m *otaServiceImpl) getRatesAvailabilitesNotAlreadyPushed(propertiesId string) ([]entity.RatesAvailabilitiesModel, string) {
	// var tableName string
	yearNow := time.Now().Year()
	monthNow := time.Now().Month()
	wg := &sync.WaitGroup{}

	roomRatesAvailabilitiesBatch := []entity.RatesAvailabilitiesModel{}
	wg.Add(12)
	for i := 0; i < 12; i++ {

		tableName := utils.GenerateTableName("rates_availabilities_", time.Now())

		go func(tableName string, wg *sync.WaitGroup) {
			roomRatesAvailabilitiesNotPushed, err := m.RatesAvailabilitiesRepository.GetRatesAvailabilitiesNotPushed(m.DB, propertiesId, tableName)
			if err != nil {
				exception.ErrorHandler(utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Get Room Rates Availabilities Not Pushed " + err.Error()})
			}

			roomRatesAvailabilitiesBatch = append(roomRatesAvailabilitiesBatch, roomRatesAvailabilitiesNotPushed...)
			wg.Done()
		}(tableName, wg)

		monthNow = monthNow + 1
		if monthNow > 12 {
			monthNow = 1
			yearNow = yearNow + 1
		}
	}

	wg.Wait()

	if len(roomRatesAvailabilitiesBatch) == 0 {
		return nil, "no data to push"
	}

	return roomRatesAvailabilitiesBatch, ""
}

func (m *otaServiceImpl) updateRatesAvailabilitesStatusPushed(ratesAvailabilitiesModel []entity.RatesAvailabilitiesModel) {
	wg := &sync.WaitGroup{}

	groupedData := make(map[string][]entity.RatesAvailabilitiesModel)

	for _, item := range ratesAvailabilitiesModel {
		groupedData[item.TableName] = append(groupedData[item.TableName], item)
	}

	wg.Add(len(groupedData))
	for tableName, data := range groupedData {
		go func(tableName string, data []entity.RatesAvailabilitiesModel, wg *sync.WaitGroup) {
			ratesAvailabilitesId := make([]string, len(data))
			for i := 0; i < len(data); i++ {
				ratesAvailabilitesId[i] = data[i].Id
			}
			m.RatesAvailabilitiesRepository.UpdateStatusRatesAvailabilitiesPushed(m.DB, ratesAvailabilitesId, tableName)
			wg.Done()
		}(tableName, data, wg)
	}
	wg.Wait()
}

func (m *otaServiceImpl) UpdateRoomsAvailabilities(c echo.Context, updateRoomsAvailabilitiesRequest *request.UpdateRoomsAvailabilitiesRequest) interface{} {

	startDateIn, _ := time.Parse("2006-01-02", updateRoomsAvailabilitiesRequest.FromDate)
	endDateOut, _ := time.Parse("2006-01-02", updateRoomsAvailabilitiesRequest.ToDate)

	roomMapping, err := m.MappingRepository.GetOtaRoomMappingByOtaRoomCode(m.DB, updateRoomsAvailabilitiesRequest.RoomCode)
	if err != nil {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Get Room Id By Room Code " + err.Error()})
	}

	if roomMapping.Id == "" {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusNotFound, Mssg: "Room Code Not Found", LogDetail: "Room Code Not Found"})
	}

	room, err := m.RoomsRepository.GetRoomsById(m.DB, roomMapping.RoomsId)
	if err != nil {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Get Room By Room Id " + err.Error()})
	}

	if room.Id == "" {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusNotFound, Mssg: "Room Id Not Found", LogDetail: "Room Id Not Found"})
	}

	ota, err := m.OtaRepository.GetOtaById(m.DB, updateRoomsAvailabilitiesRequest.OtaId)
	if err != nil {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Get Ota By Ota Id " + err.Error()})
	}

	if ota.Id == "" {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusNotFound, Mssg: "Ota Id Not Found", LogDetail: "Ota Id Not Found"})
	}

	CustomerModel := &entity.CustomerModel{}
	CustomerModel.Id = utils.RandomUUID()
	CustomerModel.UsersId = ota.Id
	CustomerModel.Name = updateRoomsAvailabilitiesRequest.CustomerName
	CustomerModel.Email = updateRoomsAvailabilitiesRequest.CustomerEmail
	CustomerModel.PhoneNumber = updateRoomsAvailabilitiesRequest.CustomerPhone
	CustomerModel.IdentityNumber = updateRoomsAvailabilitiesRequest.CustomerIdentityNumber
	CustomerModel.CreatedAt = time.Now()

	BookingModel := &entity.BookingModel{}
	BookingModel.Id = utils.RandomUUID()
	BookingModel.PropertiesId = room.Id
	BookingModel.CustomersId = CustomerModel.Id
	BookingModel.BookingCode = utils.AutoGenerateBookingCode()
	BookingModel.BookingDate = time.Now()
	BookingModel.StartDate = startDateIn
	BookingModel.EndDate = endDateOut
	BookingModel.BookingType = "ota"
	BookingModel.OtaId = updateRoomsAvailabilitiesRequest.OtaId
	BookingModel.NumberOfAdults = updateRoomsAvailabilitiesRequest.NumOfAdults
	BookingModel.NumberOfChildrens = updateRoomsAvailabilitiesRequest.NumOfChildrens
	BookingModel.TotalPrice = updateRoomsAvailabilitiesRequest.TotalPrice
	BookingModel.CreatedAt = time.Now()

	RoomReservationsModel := &entity.RoomReservationsModel{}
	RoomReservationsModel.Id = utils.RandomUUID()
	RoomReservationsModel.BookingsId = BookingModel.Id
	RoomReservationsModel.RoomsId = roomMapping.RoomsId
	RoomReservationsModel.RatePlansId = roomMapping.RatePlansId
	RoomReservationsModel.RoomPrice = 100 // total price / total night
	RoomReservationsModel.NumOfRooms = updateRoomsAvailabilitiesRequest.NumOfRooms
	RoomReservationsModel.TotalPrice = updateRoomsAvailabilitiesRequest.TotalPrice
	RoomReservationsModel.CreatedAt = time.Now()

	tx := m.DB.Begin()

	roomsAvailabilities, mssg := m.getRoomsAvailabilitesByRoomId(c, tx, roomMapping.RoomsId, startDateIn, endDateOut, updateRoomsAvailabilitiesRequest.NumOfRooms)
	if mssg != "" {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusNotFound, Mssg: mssg, LogDetail: mssg})
	}

	m.updateRoomsAvailabilities(tx, roomsAvailabilities, updateRoomsAvailabilitiesRequest.NumOfRooms)

	err = m.CustomerRepository.CreateCustomer(tx, CustomerModel)
	if err != nil {
		exception.ErrorHandlerWithRollBack(tx, utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Create Customer " + err.Error()})
	}

	err = m.BookingRepository.CreateBooking(tx, BookingModel, utils.GenerateTableName("bookings_", time.Now()))
	if err != nil {
		exception.ErrorHandlerWithRollBack(tx, utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Create Booking " + err.Error()})
	}

	err = m.RoomReservationsRepository.CreateRoomReservations(m.DB, RoomReservationsModel, utils.GenerateTableName("room_reservations_", time.Now()))
	if err != nil {
		exception.ErrorHandlerWithRollBack(tx, utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Create Room Reservations " + err.Error()})
	}

	commit := tx.Commit()

	if commit.Error != nil {
		exception.ErrorHandlerWithRollBack(tx, utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Commit Update Rooms Availabilities " + commit.Error.Error()})
	}

	return nil
}

func (m *otaServiceImpl) getRoomsAvailabilitesByRoomId(c echo.Context, tx *gorm.DB, roomId string, startDateIn, endDateOut time.Time, numOfRooms int) ([]entity.RatesAvailabilitiesModel, string) {
	roomRatesAvailabilitiesBatch := []entity.RatesAvailabilitiesModel{}
	for startDateIn.Before(endDateOut) {
		tableName := utils.GenerateTableName("rates_availabilities_", startDateIn)

		roomRatesAvailabilities, err := m.RatesAvailabilitiesRepository.GetRatesAvailabilitesByDate(c, tx, numOfRooms, roomId, startDateIn.Format("2006-01-02"), endDateOut.Format("2006-01-02"), tableName)
		if err != nil {
			exception.ErrorHandler(utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Get Room Rates Availabilities " + err.Error()})
		}

		roomRatesAvailabilitiesBatch = append(roomRatesAvailabilitiesBatch, roomRatesAvailabilities...)

		if startDateIn.Month() == 12 {
			startDateIn = time.Date(startDateIn.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
		} else {
			startDateIn = time.Date(startDateIn.Year(), startDateIn.Month()+1, 1, 0, 0, 0, 0, time.UTC)
		}
	}

	if len(roomRatesAvailabilitiesBatch) == 0 {
		return nil, "room not available"
	}

	return roomRatesAvailabilitiesBatch, ""
}

func (m *otaServiceImpl) updateRoomsAvailabilities(tx *gorm.DB, ratesAvailabilitiesModel []entity.RatesAvailabilitiesModel, numOfRooms int) {

	wg := &sync.WaitGroup{}

	groupedData := make(map[string][]entity.RatesAvailabilitiesModel)

	for _, item := range ratesAvailabilitiesModel {
		groupedData[item.TableName] = append(groupedData[item.TableName], item)
	}

	wg.Add(len(groupedData))
	errs := []error{}
	for tableName, data := range groupedData {
		go func(tableName string, data []entity.RatesAvailabilitiesModel, wg *sync.WaitGroup) {
			ratesAvailabilitesId := make([]string, len(data))
			for i := 0; i < len(data); i++ {
				ratesAvailabilitesId[i] = data[i].Id
			}
			err := m.RatesAvailabilitiesRepository.UpdateAvailabilities(tx, numOfRooms, ratesAvailabilitesId, tableName)
			if err != nil {
				errs = append(errs, err)
			}
			wg.Done()
		}(tableName, data, wg)
	}
	wg.Wait()

	if len(errs) > 0 {
		exception.ErrorHandler(utils.LogStruct{Code: http.StatusInternalServerError, Mssg: "Internal Server Error", LogDetail: "Cant Update Rooms Availabilities " + errs[0].Error()})
	}
}
