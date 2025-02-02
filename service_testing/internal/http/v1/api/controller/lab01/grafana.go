package lab01

import (
	internalErr "service_testing/internal/errors"
	"service_testing/internal/http/v1/api/schema"
	"service_testing/internal/utils"

	"github.com/labstack/echo/v4"
)

func GrafanaIsReachable(c echo.Context) error {
	ctx := c.Request().Context()
	addr := "http://localhost:3000/api/health"

	_, err := utils.RequestGetAddr(ctx, addr, 200)
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
		Message: "Grafana Is Reachable",
	})
}
