package lab01

import (
	httpSchema "service_testing/internal/http/v1/api/schema/http"
	"service_testing/internal/utils"

	"github.com/labstack/echo/v4"
)

func PrometheusTestTargetExist(c echo.Context) error {
	prometheusAddress := "http://localhost:9090"
	prometheusTargetName := "cadvisor"

	ctx := c.Request().Context()

	activeTargetsName, err := utils.GetPrometheusActiveTargetsName(ctx, prometheusAddress)
	if err != nil {
		httpError := httpSchema.HttpMessage{
			Message: err.Error(),
		}

		return c.JSON(500, httpError)
	}

	for _, targetName := range activeTargetsName {
		if targetName == prometheusTargetName {
			return c.JSON(200, httpSchema.HttpMessage{
				Message: "Target founded!",
			})
		}
	}

	return c.JSON(200, httpSchema.HttpMessage{
		Message: "Target not found!",
	})
}
