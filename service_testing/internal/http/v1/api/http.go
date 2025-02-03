package api

import (
	"service_testing/internal/http/v1/api/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	// Enable CORS for localhost:3030
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3030"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	router.SetupRouter(e)

	e.Logger.Fatal(e.Start(":1234"))
}
