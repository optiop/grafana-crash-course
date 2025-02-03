package utils

import (
	"context"
	"testing"
)

func TestGetLokiConfig(t *testing.T) {
	ctx := context.TODO()
	addr := "http://localhost:3000"

	_, err := GetLokiDatasource(ctx, addr)
	if err != nil {
		t.Fatal(err)
	}
}
