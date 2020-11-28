package gw2api

import "testing"

func TestSkins(t *testing.T) {
	api := New()
	if _, err := api.Skins(10); err != nil {
		t.Errorf("Skins failed: %s", err)
	}
}
