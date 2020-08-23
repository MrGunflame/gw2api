package util

import (
	"gitlab.com/MrGunflame/gw2api"
)

// GetLinkedWorlds returns all linked worlds of a host world in the current match
func GetLinkedWorlds(match *gw2api.WvWMatch, host int) []int {
	var clr string
	for k, v := range match.Worlds {
		if v == host {
			clr = k
			break
		}
	}

	var links []int
	for _, w := range match.AllWorlds[clr] {
		if w != host {
			links = append(links, w)
		}
	}

	return links
}

// IsLinkedWorld reports whether the entered world id is currently linked with the entered host id
func IsLinkedWorld(match *gw2api.WvWMatch, host, link int) bool {
	// Identify the host color
	var clr string
	for k, v := range match.Worlds {
		if v == host {
			clr = k
			break
		}
	}

	for _, w := range match.AllWorlds[clr] {
		if w == link {
			return true
		}
	}

	return false
}
