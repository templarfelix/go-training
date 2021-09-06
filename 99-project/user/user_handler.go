package user

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func  RegisterUserHandlers(e *echo.Echo, service *Service) {
	e.GET("/v1/users/:id", allUsers(service))
}

func allUsers(service *Service) func(echo.Context) error {
	return func(c echo.Context) error {

		//id := c.Param("id")

		user, err := service.Find(uuid.New())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, user)
	}
}
