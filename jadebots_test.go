package gw2api

import "testing"

func TestJadeBots(t *testing.T) {
	api := New()

	if _, err := api.JadeBots(); err != nil {
		t.Error(err)
	}
}
