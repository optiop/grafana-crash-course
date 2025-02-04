package utils

import (
	"context"
	"fmt"
	"net/url"
)

func GetLokiLogsExist(ctx context.Context, lokiURL string) (string, error) {
	query := `sum(rate({job="varlogs"}[10m])) by (level)`
	url := fmt.Sprintf("%s/loki/api/v1/query?query=%s", lokiURL, url.QueryEscape(query))

	res, err := RequestGetAddr(ctx, url, 200)
	if err != nil {
		return "", err
	}

	return res, nil
}
