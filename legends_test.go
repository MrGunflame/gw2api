package gw2api

import "testing"

func TestLegends(t *testing.T) {
	api := New()
	if _, err := api.Legends(); err != nil {
		t.Errorf("Legends failed: '%s'", err)
	}
}
