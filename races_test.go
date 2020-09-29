package gw2api

import "testing"

func TestRaces(t *testing.T) {
	api := New()
	if _, err := api.Races(); err != nil {
		t.Errorf("Races failed: '%s'", err)
	}
}
