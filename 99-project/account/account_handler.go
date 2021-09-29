package account

import (
	"99-project/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterAccountHandlers(e *echo.Echo, service *Service) {
	e.GET("/v1/accounts/:id", Find(service))
	e.GET("/v1/accounts", FindAll(service))
	e.POST("/v1/users", Create(service))
	e.DELETE("/v1/users/:id", Delete(service))
}

func Find(service *Service) func(echo.Context) error {
	return func(c echo.Context) error {

		id , err := entity.StringToID(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		user, err := service.Find(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, user)
	}
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

func Create(service *Service) func(echo.Context) error {

	return func(c echo.Context) error {

		u := &entity.Account{}
		if err := c.Bind(u); err != nil {
			return err
		}

		id, err := service.Create(u)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, id)
	}
}

func Delete(service *Service) func(echo.Context) error {

	return func(c echo.Context) error {

		id , err := entity.StringToID(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		err = service.Delete(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, id)
	}
}