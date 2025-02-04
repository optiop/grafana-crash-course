package utils

import "context"

func GetLokiConfiguration(ctx context.Context, lokiURL string) (string, error) {
	res, err := RequestGetAddr(ctx, lokiURL, 200, "/config")
	if err != nil {
		return "", err
	}

	return res, nil
}
