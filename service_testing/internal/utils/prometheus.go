package utils

import (
	"context"
	"errors"
	"strings"

	"github.com/prometheus/client_golang/api"
	v1Api "github.com/prometheus/client_golang/api/prometheus/v1"

	internalErr "service_testing/internal/errors"
)

func GetPrometheusActiveTargetWithName(ctx context.Context, prometheusURL, name string) (*v1Api.ActiveTarget, error) {
	client, err := api.NewClient(api.Config{
		Address: prometheusURL,
	})
	if err != nil {
		return nil, errors.Join(internalErr.ErrPrometheusCreateClient, err)
	}

	v1api := v1Api.NewAPI(client)
	targets, err := v1api.Targets(ctx)
	if err != nil {
		return nil, errors.Join(internalErr.ErrPrometheusGetAllTargets, err)
	}

	for _, target := range targets.Active {
		if jobName, ok := target.DiscoveredLabels["job"]; ok {
			if strings.EqualFold(jobName, name) {
				return &target, nil
			}
		}
	}

	return nil, internalErr.FailPrometheusNotFoundTarget
}
