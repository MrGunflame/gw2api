package chatgen

import (
	"testing"
)

func TestEncodeMapPoI(t *testing.T) {
	cases := []struct {
		given  int
		expect string
	}{
		{
			given:  56,
			expect: "[&BDgAAAA=]",
		},
	}

	for _, c := range cases {
		result := EncodeMapPoI(c.given)

		if result != c.expect {
			t.Errorf("EncodeMapPoI returned '%s' but '%s' was expected", result, c.expect)
		}
	}
}

func TestDecodeMapPoI(t *testing.T) {
	cases := []struct {
		given  string
		expect struct {
			id  int
			err error
		}
	}{
		{
			given: "[&BDgAAAA=]",
			expect: struct {
				id  int
				err error
			}{
				id: 56,
			},
		},
		{
			given: "[&BDkDAAA=]",
			expect: struct {
				id  int
				err error
			}{
				id: 825,
			},
		},
	}

	for _, c := range cases {
		result, err := DecodeMapPoI(c.given)

		if err != c.expect.err {
			t.Errorf("DecodeMapPoI returned an unexpected error: '%s'", err)
		}
		if result != c.expect.id {
			t.Errorf("DecodeMapPoI returned '%d' but '%d' was expected", result, c.expect.id)
		}
	}
}
