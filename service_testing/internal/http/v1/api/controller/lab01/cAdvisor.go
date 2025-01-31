package lab01

import (
	httpSchema "service_testing/internal/http/v1/api/schema/http"
	"service_testing/internal/utils"

	"github.com/labstack/echo/v4"
)

func CAdvisorIsReachable(c echo.Context) error {
	ctx := c.Request().Context()
	addr := "http://localhost:8081/metrics"

	_, err := utils.RequestGetAddr(ctx, addr)
	if err != nil {
		return c.JSON(500, httpSchema.HttpMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(200, httpSchema.HttpMessage{
		Message: "CAdvisor is reachable",
	})
}
