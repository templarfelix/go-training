package transaction

import (
	"99-project/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterTransactionHandlers(e *echo.Echo, service *Service) {
	e.GET("/v1/transactions/:id", Find(service))
	e.GET("/v1/transactions", FindAll(service))
	e.POST("/v1/transactions", Create(service))
}

func Find(service *Service) func(echo.Context) error {
	return func(c echo.Context) error {

		id, err := entity.StringToID(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		Transaction, err := service.Find(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, Transaction)
	}
}

func FindAll(service *Service) func(echo.Context) error {
	return func(c echo.Context) error {
		Transaction, err := service.FindAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, Transaction)
	}
}

func Create(service *Service) func(echo.Context) error {

	return func(c echo.Context) error {

		u := &entity.Transaction{}
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
