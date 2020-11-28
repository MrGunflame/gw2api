package gw2api

import "testing"

func TestNovelties(t *testing.T) {
	api := New()
	if _, err := api.Novelties(); err != nil {
		t.Errorf("Novelties failed: %s", err)
	}
}
