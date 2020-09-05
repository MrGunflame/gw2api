package chatgen

import (
	"encoding/base64"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEncodeItem(t *testing.T) {
	cases := []struct {
		given  Item
		expect string
	}{
		{
			given: Item{
				Quantity: 1,
				ID:       46762,
			},
			expect: "[&AgGqtgAA]",
		},
		{
			given: Item{
				Quantity: 1,
				ID:       46762,
				Upgrades: [2]int{24575, 0},
			},
			expect: "[&AgGqtgBA/18AAA==]",
		},
		{
			given: Item{
				Quantity: 1,
				ID:       46762,
				Upgrades: [2]int{24575, 24615},
			},
			expect: "[&AgGqtgBg/18AACdgAAA=]",
		},
		{
			given: Item{
				Quantity: 1,
				ID:       46762,
				Upgrades: [2]int{24575, 24615},
				Skin:     3709,
			},
			expect: "[&AgGqtgDgfQ4AAP9fAAAnYAAA]",
		},
	}

	for _, c := range cases {
		result := EncodeItem(c.given)

		if result != c.expect {
			t.Errorf("EncodeItem returned '%s' but '%s' was expected", result, c.expect)
		}
	}
}

func TestDecodeItem(t *testing.T) {
	cases := []struct {
		given  string
		expect struct {
			item *Item
			err  error
		}
	}{
		{
			given: "?",
			expect: struct {
				item *Item
				err  error
			}{
				item: nil,
				err:  base64.CorruptInputError(0),
			},
		},
		{
			given: "[&AgGq]",
			expect: struct {
				item *Item
				err  error
			}{
				err: ErrInvalidHeader,
			},
		},
		{
			given: "[&AgGqtgAA]",
			expect: struct {
				item *Item
				err  error
			}{
				item: &Item{
					Quantity: 1,
					ID:       46762,
				},
				err: nil,
			},
		},
		{
			given: "[&AgGqtgBA/18AAA==]",
			expect: struct {
				item *Item
				err  error
			}{
				item: &Item{
					Quantity: 1,
					ID:       46762,
					Upgrades: [2]int{24575, 0},
				},
				err: nil,
			},
		},
		{
			given: "[&AgGqtgBg/18AACdgAAA=]",
			expect: struct {
				item *Item
				err  error
			}{
				item: &Item{
					Quantity: 1,
					ID:       46762,
					Upgrades: [2]int{24575, 24615},
				},
				err: nil,
			},
		},
		{
			given: "[&AgGqtgCAfQ4AAA==]",
			expect: struct {
				item *Item
				err  error
			}{
				item: &Item{
					Quantity: 1,
					ID:       46762,
					Skin:     3709,
				},
				err: nil,
			},
		},
		{
			given: "[&AgGqtgDAfQ4AAP9fAAA=]",
			expect: struct {
				item *Item
				err  error
			}{
				item: &Item{
					Quantity: 1,
					ID:       46762,
					Upgrades: [2]int{24575, 0},
					Skin:     3709,
				},
				err: nil,
			},
		},
		{
			given: "[&AgGqtgDgfQ4AAP9fAAAnYAAA]",
			expect: struct {
				item *Item
				err  error
			}{
				item: &Item{
					Quantity: 1,
					ID:       46762,
					Upgrades: [2]int{24575, 24615},
					Skin:     3709,
				},
				err: nil,
			},
		},
	}

	for _, c := range cases {
		result, err := DecodeItem(c.given)

		if err != c.expect.err {
			t.Errorf("DecodeItem returned an unexpected error: '%s'", err)
		}
		if !cmp.Equal(result, c.expect.item) {
			t.Errorf("DecodeItem returned unexpected data: '%s'", cmp.Diff(result, c.expect.item))
		}
	}
}
