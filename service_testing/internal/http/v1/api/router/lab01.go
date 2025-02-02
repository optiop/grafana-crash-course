package router

import (
	"service_testing/internal/http/v1/api/controller/lab01"

	"github.com/labstack/echo/v4"
)

func setupRouterLab01(g *echo.Group) {
	rg := g.Group("/lab01")

	rg.GET("/", lab01.ServiceLab01)
	rg.GET("/healthy", lab01.Healthy)

	// localhots:1234/api/v1/lab01/test_grafana_reachable

	// test_
	rg.GET("/test_grafana_reachable", lab01.GrafanaIsReachable)
	rg.GET("/test_cadvisor_docker_up", lab01.CAdvisorDockerIsUp)
	rg.GET("/test_cadvisor_reachable", lab01.CAdvisorIsReachable)
	rg.GET("/test_prometheus_target_exist", lab01.PrometheusTestTargetExist)
}
