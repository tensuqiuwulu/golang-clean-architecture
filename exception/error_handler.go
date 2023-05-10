package exception

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/response"
	"github.com/tensuqiuwulu/golang-clean-architecture/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ErrorHandler(logStruct utils.LogStruct) {
	out, err := json.Marshal(logStruct)

	if err != nil {
		panic(err)
	}

	panic(string(out))
}

func ErrorHandlerWithRollBack(db *gorm.DB, logStruct utils.LogStruct) {
	log.Println("Rollback")
	rollback := db.Rollback()
	if rollback.Error != nil {
		panic(rollback.Error)
	}

	out, err := json.Marshal(logStruct)

	if err != nil {
		panic(err)
	}

	panic(string(out))
}

func MakeHTTPErrorHandler(err error, c echo.Context) {
	logStruct := utils.LogStruct{}
	json.Unmarshal([]byte(err.Error()), &logStruct)
	utils.MakeLogEntry(c).Errorf("requestID: %s, error: %s", c.Response().Header().Get(echo.HeaderXRequestID), err.Error())
	if logStruct.Code != 0 {
		response := response.ApiResponseModel{Code: logStruct.Code, Mssg: logStruct.Mssg, Data: logStruct.Data}
		c.JSON(logStruct.Code, response)
	} else {
		response := response.ApiResponseModel{Code: logStruct.Code, Mssg: "Internal Server Error", Data: logStruct.Data}
		c.JSON(http.StatusInternalServerError, response)
	}
}
