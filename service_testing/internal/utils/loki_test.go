package utils

import (
	"context"
	"testing"
)

func TestGetLokiConfig(t *testing.T) {
	ctx := context.TODO()
	addr := "http://localhost:9080"

	_, err := GetPromtailConfiguration(ctx, addr)
	if err != nil {
		t.Fatal(err)
	}
}
