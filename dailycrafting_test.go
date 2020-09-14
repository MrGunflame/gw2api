package gw2api

import "testing"

func TestDailyCrafting(t *testing.T) {
	api := New()
	if _, err := api.DailyCrafting(); err != nil {
		t.Errorf("DailyCrafting failed: '%s'", err)
	}
}
