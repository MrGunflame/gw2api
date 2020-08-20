package gw2api

import "testing"

func TestItemStats(t *testing.T) {
	api := New()

	if _, err := api.GetItemStats(); err != nil {
		t.Error(err)
	}
}
