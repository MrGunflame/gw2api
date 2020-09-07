package gw2api

import (
	"testing"
)

func TestTile(t *testing.T) {
	api := New()
	if _, err := api.Tile(2, 1, 2, 2, 3); err != nil {
		t.Errorf("Tile failed fetching endpoint: '%s'", err)
	}
}
