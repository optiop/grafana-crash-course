package utils

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/api"
	v1Api "github.com/prometheus/client_golang/api/prometheus/v1"
)

func GetPrometheusActiveTargetsName(ctx context.Context, prometheusURL string) ([]string, error) {
	client, err := api.NewClient(api.Config{
		Address: prometheusURL,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating Prometheus client: %w", err)
	}

	v1api := v1Api.NewAPI(client)
	targets, err := v1api.Targets(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting targets: %w", err)
	}

	var targetNames []string
	for _, target := range targets.Active {
		targetNames = append(targetNames, target.DiscoveredLabels["job"])
	}

	return targetNames, nil
}
