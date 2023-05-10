package api

import (
	"github.com/tensuqiuwulu/golang-clean-architecture/src/handler"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/middleware"

	"github.com/labstack/echo/v4"
)

func OtaRoute(e *echo.Echo, otaHandler handler.OtaHandler, middleware middleware.Middleware) {
	group := e.Group("api/v1")
	group.POST("/ota/pull-rooms-rates", otaHandler.FetchRoomsAndRatesFromOta, middleware.CheckInternalToken)
	group.POST("/ota/push-rates-availabilities", otaHandler.PushRoomsAndRatesToOta, middleware.CheckInternalToken)
	group.POST("/ota/update-rooms-availabilities", otaHandler.UpdateRoomsAvailabilities)
}

func MainRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "Status OK!")
	})
}
