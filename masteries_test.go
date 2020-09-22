package gw2api

import "testing"

func TestMasteries(t *testing.T) {
	api := New()
	if _, err := api.Masteries(); err != nil {
		t.Errorf("Masteries failed: '%s'", err)
	}
}
