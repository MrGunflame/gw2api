package gw2api

import "testing"

func TestMountSkins(t *testing.T) {
	api := New()
	if _, err := api.MountSkins(); err != nil {
		t.Errorf("MountSkins failed: '%s'", err)
	}
}

func TestMountTypes(t *testing.T) {
	api := New()
	if _, err := api.MountTypes(); err != nil {
		t.Errorf("MountTypes failed: '%s'", err)
	}
}
