package gw2api

import "testing"

func TestOutfits(t *testing.T) {
	api := New()
	if _, err := api.Outfits(); err != nil {
		t.Errorf("Outfits failed: '%s'", err)
	}
}
