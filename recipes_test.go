package gw2api

import "testing"

func TestRecipes(t *testing.T) {
	api := New()
	if _, err := api.Recipes(7319); err != nil {
		t.Errorf("Recipes failed: %s", err)
	}
}
