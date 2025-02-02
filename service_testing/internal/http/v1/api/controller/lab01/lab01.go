package lab01

import (
	"service_testing/internal/http/v1/api/schema"

	"github.com/labstack/echo/v4"
)

func ServiceLab01(c echo.Context) error {
	message := schema.Message{
		Message: "Lab 01 Service Is Healthy",
	}

	return c.JSON(200, message)
}

func Healthy(c echo.Context) error {
	message := schema.Message{
		Message: "Lab 01 Service Is Healthy",
	}

	return c.JSON(200, message)
}
