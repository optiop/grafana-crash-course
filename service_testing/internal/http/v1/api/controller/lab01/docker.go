package lab01

import (
	httpSchema "service_testing/internal/http/v1/api/schema/http"
	"service_testing/internal/utils"

	"github.com/labstack/echo/v4"
)

func CAdvisorDockerIsUp(c echo.Context) error {
	ctx := c.Request().Context()
	containerName := "lab.01-cadvisor"

	err := utils.GetDockerContainerInfo(ctx, containerName)
	if err != nil {
		return c.JSON(500, httpSchema.HttpMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(200, httpSchema.HttpMessage{
		Message: "Container CAdvisor is up",
	})
}
