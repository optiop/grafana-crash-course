package lab01

import (
	httpSchema "service_testing/internal/http/v1/api/schema/http"

	"github.com/labstack/echo/v4"
)

func ServiceLab01(c echo.Context) error {
	message := httpSchema.HttpMessage{
		Message: "Lav 01 Service Is Healthy",
	}

	return c.JSON(200, message)
}

func Healthy(c echo.Context) error {
	message := httpSchema.HttpMessage{
		Message: "Lav 01 Service Is Healthy",
	}

	return c.JSON(200, message)
}
