package gw2api

import "testing"

func TestRaids(t *testing.T) {
	api := New()
	if _, err := api.Raids(); err != nil {
		t.Errorf("Raids failed: %s", err)
	}
}
