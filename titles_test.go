package gw2api

import "testing"

func TestTitles(t *testing.T) {
	api := New()
	if _, err := api.Titles(); err != nil {
		t.Errorf("Titles failed: %s", err)
	}
}
