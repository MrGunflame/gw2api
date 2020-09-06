package chatgen

import "testing"

func TestEncodeRecipe(t *testing.T) {
	cases := []struct {
		given  int
		expect string
	}{
		{
			given:  1,
			expect: "[&CQEAAAA=]",
		},
		{
			given:  2,
			expect: "[&CQIAAAA=]",
		},
		{
			given:  7,
			expect: "[&CQcAAAA=]",
		},
	}

	for _, c := range cases {
		result := EncodeRecipe(c.given)

		if result != c.expect {
			t.Errorf("EncodeRecipe returned '%s' but '%s' was expected", result, c.expect)
		}
	}
}

func TestDecodeRecipe(t *testing.T) {
	cases := []struct {
		given  string
		expect struct {
			id  int
			err error
		}
	}{
		{
			given: "[&CQEAAAA=]",
			expect: struct {
				id  int
				err error
			}{
				id: 1,
			},
		},
	}

	for _, c := range cases {
		result, err := DecodeRecipe(c.given)

		if err != c.expect.err {
			t.Errorf("DecodeRecipe returned an unexpected error: '%s'", err)
		}
		if result != c.expect.id {
			t.Errorf("DecodeRecipe returned '%d' but '%d' was expected", result, c.expect.id)
		}
	}
}
