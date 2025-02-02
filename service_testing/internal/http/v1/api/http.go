package api

import (
	"service_testing/internal/http/v1/api/router"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	router.SetupRouter(e)

	e.Logger.Fatal(e.Start(":1234"))
}
