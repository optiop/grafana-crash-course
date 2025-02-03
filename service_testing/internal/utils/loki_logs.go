package utils

import "context"

func GetLokiLogsExist(ctx context.Context, lokiURL string) (string, error) {
	res, err := RequestGetAddr(ctx, lokiURL, 200, `/loki/api/v1/query --data-urlencode 'query=sum(rate({job="varlogs"}[10m])) by (level)' | jq`)
	if err != nil {
		return "", err
	}

	return res, nil
}
