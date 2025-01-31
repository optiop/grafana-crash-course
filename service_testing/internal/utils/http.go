package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func RequestGetAddr(ctx context.Context, addr string, path ...string) (string, error) {
	addr, err := url.JoinPath(addr, path...)
	if err != nil {
		return "", err
	}

	httpClient := http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, addr, nil)
	if err != nil {
		return "", err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
