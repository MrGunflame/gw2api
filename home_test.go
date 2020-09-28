package gw2api

import "testing"

func TestHomeCats(t *testing.T) {
	api := New()
	if _, err := api.HomeCats(); err != nil {
		t.Errorf("HomeCats failed: '%s'", err)
	}
}

func TestHomeNodes(t *testing.T) {
	api := New()
	if _, err := api.HomeNodes(); err != nil {
		t.Errorf("HomeNodes failed: '%s'", err)
	}
}
