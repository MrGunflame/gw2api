package gw2api

import "testing"

func TestStories(t *testing.T) {
	api := New()
	if _, err := api.Stories(); err != nil {
		t.Errorf("Stories failed: %s", err)
	}
}
