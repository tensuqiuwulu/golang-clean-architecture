package traveloka

import "github.com/labstack/echo/v4"

type TravelokaDomainController interface {
	PushRoomsRates(c echo.Context) error
}

type travelokaDomainImpl struct {
}

func NewTravelokaDomain() TravelokaDomainController {
	return &travelokaDomainImpl{}
}

func (t *travelokaDomainImpl) PushRoomsRates(c echo.Context) error {
	return nil
}
