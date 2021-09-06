package user

import (
	"99-project/entity"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterUserHandlers(e *echo.Echo, service *Service) {
	e.GET("/v1/users/:id", Find(service))
	e.GET("/v1/users", FindAll(service))
	e.POST("/v1/users", Store(service))
}

func Find(service *Service) func(echo.Context) error {
	return func(c echo.Context) error {

		//id := c.Param("id")

		user, err := service.Find(uuid.New())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, user)
	}
}

func FindAll(service *Service) func(echo.Context) error {
	return func(c echo.Context) error {
		user, err := service.FindAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, user)
	}
}

func Store(service *Service) func(echo.Context) error {

	return func(c echo.Context) error {

		u := &entity.User{}
		if err := c.Bind(u); err != nil {
			return err
		}

		id, err := service.Store(u)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, id)
	}
}
