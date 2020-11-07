package util

import (
	"fmt"
	"math"

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

// PredictNextWvWMatches predicts the next weeks matchup based on the current point standings
// The lowest tier cannot descend, the highest cannot ascend
// The returned match only returns id and word data
func PredictNextWvWMatches(matches []*gw2api.WvWMatch) []*gw2api.WvWMatch {
	nextMatches := make([]*gw2api.WvWMatch, len(matches))
	for i := range nextMatches {
		nextMatches[i] = &gw2api.WvWMatch{
			ID:     matches[i].ID,
			Worlds: make(map[string]int),
		}
	}

	inSlice := func(sl []string, s string) bool {
		for _, v := range sl {
			if v == s {
				return true
			}
		}
		return false
	}

	notIn := func(m *map[string]int, clr []string) string {
		for k := range *m {
			if !inSlice(clr, k) {
				return k
			}
		}
		return ""
	}

	for i, m := range matches {
		fmt.Println(i)

		var low, high string
		lowVal, highVal := math.MaxInt64, 0
		for k, v := range m.VictoryPoints {
			if v < lowVal {
				low = k
				lowVal = v
			}
			if v > highVal {
				high = k
				highVal = v
			}

		}

		// Top tier - cannot ascend further
		if i == 0 {
			nextMatches[i].Worlds["green"] = m.Worlds[high]
		} else {
			nextMatches[i-1].Worlds["red"] = m.Worlds[high]
		}

		// Bottom tier - cannot descend further
		if i == len(matches)-1 {
			nextMatches[i].Worlds["red"] = m.Worlds[low]
		} else {
			nextMatches[i+1].Worlds["green"] = m.Worlds[low]
		}

		nextMatches[i].Worlds["blue"] = m.Worlds[notIn(&m.Worlds, []string{low, high})]
	}

	return nextMatches
}
