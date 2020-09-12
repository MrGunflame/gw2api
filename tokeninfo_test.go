package gw2api

import (
	"os"
	"testing"
)

func TestTokeninfo(t *testing.T) {
	key := os.Getenv("APIKEY")
	if key == "" {
		t.Skip("Unable to test without APIKEY")
	}

	api := New().WithAccessToken(key)
	if _, err := api.Tokeninfo(); err != nil {
		t.Errorf("Tokeninfo failed: '%s'", err)
	}
}
