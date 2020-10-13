package gw2api

import "testing"

func TestItems(t *testing.T) {
	api := New()
	if _, err := api.Items(30704, 67528, 48085, 24609, 67339); err != nil {
		t.Errorf("Items failed: '%s'", err)
	}
}
