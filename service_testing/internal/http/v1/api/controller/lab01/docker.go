package lab01

import (
	internalErr "service_testing/internal/errors"
	"service_testing/internal/http/v1/api/schema"
	"service_testing/internal/utils"

	"github.com/labstack/echo/v4"
)

func CAdvisorDockerIsUp(c echo.Context) error {
	ctx := c.Request().Context()
	containerName := "lab.01-cadvisor"

	err := utils.GetDockerContainerInfo(ctx, containerName)
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
		Message: "Container CAdvisor is up",
	})
}
