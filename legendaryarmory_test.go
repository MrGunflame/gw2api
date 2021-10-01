package gw2api

import "testing"

func TestLegendaryArmory(t *testing.T) {
	api := New()
	if _, err := api.LegendaryArmory(); err != nil {
		t.Errorf("LegendaryArmory failed: '%s'", err)
	}
}
