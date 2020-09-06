package chatgen

import "testing"

func TestEncodeTrait(t *testing.T) {
	cases := []struct {
		given  int
		expect string
	}{
		{
			given:  1010,
			expect: "[&B/IDAAA=]",
		},
	}

	for _, c := range cases {
		result := EncodeTrait(c.given)

		if result != c.expect {
			t.Errorf("EncodeTrait returned '%s' but '%s' was expected", result, c.expect)
		}
	}
}

func TestDecodeTrait(t *testing.T) {
	cases := []struct {
		given  string
		expect struct {
			id  int
			err error
		}
	}{
		{
			given: "[&B/IDAAA=]",
			expect: struct {
				id  int
				err error
			}{
				id: 1010,
			},
		},
	}

	for _, c := range cases {
		result, err := DecodeTrait(c.given)

		if err != c.expect.err {
			t.Errorf("DecodeTrait returned an unexpected error: '%s'", err)
		}
		if result != c.expect.id {
			t.Errorf("DecodeTrait returned '%d' but '%d' was expected", result, c.expect.id)
		}
	}
}
