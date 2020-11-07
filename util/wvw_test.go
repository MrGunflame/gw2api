package util

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"gitlab.com/MrGunflame/gw2api"
)

func TestPredictNextWvWMatches(t *testing.T) {
	cases := []struct {
		given  []*gw2api.WvWMatch
		expect []*gw2api.WvWMatch
	}{
		{
			given: []*gw2api.WvWMatch{
				{
					ID: "2-1",
					Worlds: map[string]int{
						"red":   2013,
						"blue":  2206,
						"green": 2010,
					},
					AllWorlds: map[string][]int{
						"red":   []int{2009, 2013},
						"blue":  []int{2204, 2206},
						"green": []int{2008, 2010},
					},
					VictoryPoints: map[string]int{
						"red":   78,
						"blue":  98,
						"green": 109,
					},
				},
				{
					ID: "2-2",
					Worlds: map[string]int{
						"red":   2203,
						"blue":  2301,
						"green": 2202,
					},
					AllWorlds: map[string][]int{
						"red":   []int{2006, 2203},
						"blue":  []int{2301},
						"green": []int{2207, 2202},
					},
					VictoryPoints: map[string]int{
						"red":   86,
						"blue":  88,
						"green": 102,
					},
				},
				{
					ID: "2-3",
					Worlds: map[string]int{
						"red":   2014,
						"blue":  2007,
						"green": 2012,
					},
					AllWorlds: map[string][]int{
						"red":   []int{2201, 2014},
						"blue":  []int{2011, 2007},
						"green": []int{2004, 2012},
					},
					VictoryPoints: map[string]int{
						"red":   79,
						"blue":  94,
						"green": 103,
					},
				},
			},
			expect: []*gw2api.WvWMatch{
				{
					ID: "2-1",
					Worlds: map[string]int{
						"red":   2202,
						"blue":  2206,
						"green": 2010,
					},
					AllWorlds: map[string][]int{
						"red":   []int{2207, 2202},
						"blue":  []int{2204, 2206},
						"green": []int{2008, 2010},
					},
				},
				{
					ID: "2-2",
					Worlds: map[string]int{
						"red":   2012,
						"blue":  2301,
						"green": 2013,
					},
					AllWorlds: map[string][]int{
						"red":   []int{2004, 2012},
						"blue":  []int{2301},
						"green": []int{2009, 2013},
					},
				},
				{
					ID: "2-3",
					Worlds: map[string]int{
						"red":   2014,
						"blue":  2007,
						"green": 2203,
					},
					AllWorlds: map[string][]int{
						"red":   []int{2201, 2014},
						"blue":  []int{2011, 2007},
						"green": []int{2006, 2203},
					},
				},
			},
		},
	}

	for _, c := range cases {
		result := PredictNextWvWMatches(c.given)
		if !cmp.Equal(result, c.expect) {
			t.Errorf("PredictNextWvWMatches failed: '%s'", cmp.Diff(result, c.expect))
		}
	}
}
