package chatgen

import "testing"

func TestEncodeOutfit(t *testing.T) {
	cases := []struct {
		given  int
		expect string
	}{
		{
			given:  4,
			expect: "[&CwQAAAA=]",
		},
	}

	for _, c := range cases {
		result := EncodeOutfit(c.given)

		if result != c.expect {
			t.Errorf("EncodeOutfit returned '%s' but '%s' was expected", result, c.expect)
		}
	}
}

func TestDecodeOutfit(t *testing.T) {
	cases := []struct {
		given  string
		expect struct {
			id  int
			err error
		}
	}{
		{
			given: "[&CwQAAAA=]",
			expect: struct {
				id  int
				err error
			}{
				id: 4,
			},
		},
	}

	for _, c := range cases {
		result, err := DecodeOutfit(c.given)

		if err != c.expect.err {
			t.Errorf("DecodeOutfit returned an unexpected error: '%s'", err)
		}
		if result != c.expect.id {
			t.Errorf("DecodeOutfit returned '%d' but '%d' was expected", result, c.expect.id)
		}
	}
}
