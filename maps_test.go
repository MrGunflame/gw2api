package gw2api

import "testing"

func TestMaps(t *testing.T) {
	api := New()
	if _, err := api.Maps(); err != nil {
		t.Errorf("Maps failed: %s", err)
	}
}
