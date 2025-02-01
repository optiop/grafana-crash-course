package utils

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"

	internalErr "service_testing/internal/errors"
)

func RequestGetAddr(ctx context.Context, addr string, statusCode int, path ...string) (string, error) {
	addr, err := url.JoinPath(addr, path...)
	if err != nil {
		return "", errors.Join(internalErr.ErrHttpInvalidUrl, err)
	}

	httpClient := http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, addr, nil)
	if err != nil {
		return "", errors.Join(internalErr.ErrHttpCreateNewReq, err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) {
			return "", errors.Join(internalErr.FailHttpUnreachable, err)
		}
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != statusCode {
		return "", internalErr.ErrHttpNotValidStatus
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Join(internalErr.ErrHttpParsBody, err)
	}

	return string(data), nil
}
