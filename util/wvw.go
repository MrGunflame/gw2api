package util

import (
	"gitlab.com/MrGunflame/gw2api"
)

// IsLinkedWorld reports whether the entered world id is currently linked with the entered host id
func IsLinkedWorld(match *gw2api.WvWMatch, host, link int) bool {
	// Identify the host color
	var clr string
	for k, v := range match.Worlds {
		if v == host {
			clr = k
		}
	}

	for _, w := range match.AllWorlds[clr] {
		if w == link {
			return true
		}
	}

	return false
}
