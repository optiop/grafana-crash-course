package lab01

import (
	"service_testing/internal/http/v1/api/schema"
	"service_testing/internal/utils"

	internalErr "service_testing/internal/errors"

	"github.com/labstack/echo/v4"
)

func PrometheusTestTargetExist(c echo.Context) error {
	prometheusAddress := "http://localhost:9090"
	prometheusTargetName := "cadvisor"

	ctx := c.Request().Context()

	_, err := utils.GetPrometheusActiveTargetWithName(ctx, prometheusAddress, prometheusTargetName)
	if err != nil {
		if errTypeName := internalErr.TypeErrorDetection(err); errTypeName != "" {
			return c.JSON(200, schema.Message{
				Message: err.Error(),
				Status:  errTypeName,
			})
		}

		return c.JSON(500, schema.Message{
			Message: err.Error(),
			Status:  "panic",
		})
	}

	return c.JSON(200, schema.Message{
		Message: "Target founded!",
	})
}
