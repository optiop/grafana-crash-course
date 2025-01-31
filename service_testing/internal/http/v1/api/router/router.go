package router

import "github.com/labstack/echo/v4"

func SetupRouter(e *echo.Echo) {
	g := e.Group("/api/v1")

	setupRouterLab01(g)
}
