package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/tensuqiuwulu/golang-clean-architecture/exception"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/request"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/response"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/service"
	"github.com/tensuqiuwulu/golang-clean-architecture/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type OtaHandler interface {
	FetchRoomsAndRatesFromOta(c echo.Context) error
	PushRoomsAndRatesToOta(c echo.Context) error
	UpdateRoomsAvailabilities(c echo.Context) error
}

type otaHandlerImpl struct {
	Validate   *validator.Validate
	OtaService service.OtaService
}

func NewOtaHandler(
	validate *validator.Validate,
	otaService service.OtaService,
) OtaHandler {
	return &otaHandlerImpl{
		Validate:   validate,
		OtaService: otaService,
	}
}

func (m *otaHandlerImpl) FetchRoomsAndRatesFromOta(c echo.Context) error {
	var err error

	pullRoomsAndRatesFromOtaRequest := &request.PullRoomsAndRatesFromOtaRequest{}
	if err = c.Bind(pullRoomsAndRatesFromOtaRequest); err != nil {
		exception.ErrorHandler(utils.LogStruct{Code: 400, Mssg: "Bad Request", LogDetail: err.Error()})
	}

	err, detail := utils.ValidateRequest(m.Validate, pullRoomsAndRatesFromOtaRequest)
	if err != nil {
		exception.ErrorHandler(utils.LogStruct{Code: 400, Mssg: "Bad Request", LogDetail: "Validation Error", Data: detail})
	}

	respData := m.OtaService.FetchRoomsAndRatesFromOta(c, pullRoomsAndRatesFromOtaRequest)
	response := response.ApiResponseModel{Code: http.StatusOK, Mssg: "success", Data: respData}
	return c.JSON(http.StatusOK, response)
}

func (m *otaHandlerImpl) PushRoomsAndRatesToOta(c echo.Context) error {
	var err error

	pushRoomsAndRatesToOtaRequest := &request.PushRoomsAndRatesToOtaRequest{}
	if err = c.Bind(pushRoomsAndRatesToOtaRequest); err != nil {
		exception.ErrorHandler(utils.LogStruct{Code: 500, Mssg: "Internal Server Error", LogDetail: err.Error()})
	}

	err, detail := utils.ValidateRequest(m.Validate, pushRoomsAndRatesToOtaRequest)
	if err != nil {
		exception.ErrorHandler(utils.LogStruct{Code: 400, Mssg: "Bad Request", LogDetail: "Validation Error", Data: detail})
	}

	respData := m.OtaService.PushRoomsAndRatesToOta(c, pushRoomsAndRatesToOtaRequest)
	response := response.ApiResponseModel{Code: http.StatusOK, Mssg: "success", Data: respData}
	return c.JSON(http.StatusOK, response)
}

func (m *otaHandlerImpl) UpdateRoomsAvailabilities(c echo.Context) error {
	var err error

	updateRoomsAvailabilitiesRequest := &request.UpdateRoomsAvailabilitiesRequest{}
	start := time.Now()
	if err = c.Bind(updateRoomsAvailabilitiesRequest); err != nil {
		exception.ErrorHandler(utils.LogStruct{Code: 500, Mssg: "Internal Server Error", LogDetail: err.Error()})
	}

	err, detail := utils.ValidateRequest(m.Validate, updateRoomsAvailabilitiesRequest)
	if err != nil {
		exception.ErrorHandler(utils.LogStruct{Code: 400, Mssg: "Bad Request", LogDetail: "Validation Error", Data: detail})
	}

	resp := m.OtaService.UpdateRoomsAvailabilities(c, updateRoomsAvailabilitiesRequest)
	response := response.ApiResponseModel{Code: http.StatusOK, Mssg: "success", Data: resp}
	log.Println("took = ", time.Since(start))
	return c.JSON(http.StatusOK, response)
}
