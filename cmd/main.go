package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/tensuqiuwulu/golang-clean-architecture/config"
	"github.com/tensuqiuwulu/golang-clean-architecture/config/database"
	"github.com/tensuqiuwulu/golang-clean-architecture/exception"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/api"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/handler"
	appMiddleware "github.com/tensuqiuwulu/golang-clean-architecture/src/middleware"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/repository"
	"github.com/tensuqiuwulu/golang-clean-architecture/src/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Get Config
	appConfig := config.GetConfig()

	// Mysql Connection
	DBConn := database.NewDBConnection(&appConfig.Database)

	validate := validator.New()

	// Timezone
	location, err := time.LoadLocation(appConfig.Timezone.Timezone)
	log.Println("Location:", location, err)

	// Server App
	log.Println("Server App : ", string(appConfig.Application.Server))

	e := echo.New()

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisableStackAll:   true,
		DisablePrintStack: true,
	}))

	e.HTTPErrorHandler = exception.MakeHTTPErrorHandler
	e.Use(middleware.RequestID())

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Second,
	}))

	bookingRepository := repository.NewBookingRepository()
	roomReservationsRepository := repository.NewRoomReservationsRepository()
	customerRepository := repository.NewCustomerRepository()
	ratesAvailabilitiesRepository := repository.NewRatesAvailabilitiesRepository()
	otaRepository := repository.NewOtaRepository()
	mappingRepository := repository.NewMappingRepository()
	otaApiRepository := repository.NewOtaApiRepository()
	oauthRepository := repository.NewOauthRepository()
	roomsRepository := repository.NewRoomsRepository()

	otaService := service.NewOtaService(DBConn, mappingRepository, otaApiRepository, otaRepository, ratesAvailabilitiesRepository, customerRepository, bookingRepository, roomReservationsRepository, roomsRepository)

	otaHandler := handler.NewOtaHandler(validate, otaService)

	middleware := appMiddleware.NewMiddleware(DBConn, oauthRepository)
	api.OtaRoute(e, otaHandler, middleware)
	api.MainRoute(e)

	// Careful shutdown
	go func() {
		if err := e.Start(":" + strconv.Itoa(int(appConfig.Webserver.Port))); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	log.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// mysql database
	database.DBClose(DBConn)
	log.Println("Echo was successful shutdown.")
}
