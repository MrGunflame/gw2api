package gw2api

import "testing"

func TestContinents(t *testing.T) {
	api := New()
	if _, err := api.Continents(); err != nil {
		t.Errorf("Continents failed: %s", err)
	}
}
