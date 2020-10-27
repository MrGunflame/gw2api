package gw2api

import "testing"

func TestWorld(t *testing.T) {
	api := New()

	if _, err := api.Worlds(); err != nil {
		t.Error("GetWorld failed: ", err)
	}
}
