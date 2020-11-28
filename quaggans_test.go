package gw2api

import "testing"

func TestQuaggans(t *testing.T) {
	api := New()
	if _, err := api.Quaggans(); err != nil {
		t.Errorf("Quaggans failed: %s", err)
	}
}
