package gw2api

import "testing"

func TestCurrencies(t *testing.T) {
	api := New()
	if _, err := api.Currencies(); err != nil {
		t.Errorf("Currencies failed: %s", err)
	}
}
