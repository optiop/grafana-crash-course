package router

import (
	"service_testing/internal/http/v1/api/controller/lab03"

	"github.com/labstack/echo/v4"
)

func setupRouterLab03(g *echo.Group) {
	rg := g.Group("/lab03")

	rg.GET("/test_loki_configuration_not_found", lab03.LokiConfigurationNotFound)
	rg.GET("/test_grafana_loki_datasource_not_found", lab03.GrafanaLokiDatasourceNotFound)
	rg.GET("/test_promtail_configuration_not_found", lab03.PromtailConfigurationNotFound)
	rg.GET("/test_loki_logs_not_found", lab03.LokiLogsNotFound)
}
