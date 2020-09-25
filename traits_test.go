package gw2api

import "testing"

func TestTraits(t *testing.T) {
	api := New()
	if _, err := api.Traits(); err != nil {
		t.Errorf("Traits failed: '%s'", err)
	}
}
