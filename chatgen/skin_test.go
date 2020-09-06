package chatgen

import "testing"

func TestEncodeSkin(t *testing.T) {
	cases := []struct {
		given  int
		expect string
	}{
		{
			given:  4,
			expect: "[&CgQAAAA=]",
		},
	}

	for _, c := range cases {
		result := EncodeSkin(c.given)

		if result != c.expect {
			t.Errorf("EncodeSkin returned '%s' but '%s' was expected", result, c.expect)
		}
	}
}

func TestDecodeSkin(t *testing.T) {
	cases := []struct {
		given  string
		expect struct {
			id  int
			err error
		}
	}{
		{
			given: "[&CgQAAAA=]",
			expect: struct {
				id  int
				err error
			}{
				id: 4,
			},
		},
	}

	for _, c := range cases {
		result, err := DecodeSkin(c.given)

		if err != c.expect.err {
			t.Errorf("DecodeSkin returned an unexpected error: '%s'", err)
		}
		if result != c.expect.id {
			t.Errorf("DecodeSkin returned '%d' but '%d' was expected", result, c.expect.id)
		}
	}
}
