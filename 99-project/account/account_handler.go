package account

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterAccountHandlers(e *echo.Echo, service *Service) {
	e.GET("/v1/accounts", FindAll(service))
}

func FindAll(service *Service) func(echo.Context) error {
	return func(c echo.Context) error {
		Account, err := service.FindAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, Account)
	}
}