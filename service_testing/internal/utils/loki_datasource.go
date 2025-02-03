package utils

import "context"

func GetLokiDatasource(ctx context.Context, lokiURL string) (string, error) {
	res, err := RequestGetAddr(ctx, lokiURL, 200, "/api/datasources")
	if err != nil {
		return "", err
	}

	return res, nil
}
