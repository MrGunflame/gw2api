package gw2api

import "testing"

func TestColors(t *testing.T) {
	api := New()
	if _, err := api.Colors(); err != nil {
		t.Errorf("Colors failed: %s", err)
	}
}
