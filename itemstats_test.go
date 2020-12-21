package gw2api

import "testing"

func TestItemStats(t *testing.T) {
	api := New()

	if _, err := api.ItemStats(); err != nil {
		t.Error(err)
	}
}
