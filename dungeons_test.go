package gw2api

import "testing"

func TestDungeons(t *testing.T) {
	api := New()
	if _, err := api.Dungeons(); err != nil {
		t.Errorf("Dungeons failed: %s", err)
	}
}
