package gw2api

import "testing"

func TestMaterials(t *testing.T) {
	api := New()
	if _, err := api.Materials(); err != nil {
		t.Errorf("Materials failed: '%s'", err)
	}
}
