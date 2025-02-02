package utils

import (
	"context"
	"testing"
)

func TestGetLokiConfig(t *testing.T) {
	ctx := context.TODO()
	addr := "http://localhost:3100"

	_, err := GetLokiConfiguration(ctx, addr)
	if err != nil {
		t.Fatal(err)
	}
}
