package gw2api

import "testing"

func TestPets(t *testing.T) {
	api := New()
	if _, err := api.Pets(); err != nil {
		t.Errorf("Pets failed: '%s'", err)
	}
}
