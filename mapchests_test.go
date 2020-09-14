package gw2api

import "testing"

func TestMapChests(t *testing.T) {
	api := New()
	if _, err := api.MapChests(); err != nil {
		t.Errorf("MapChests failed: '%s'", err)
	}
}
