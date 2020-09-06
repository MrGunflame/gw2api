package chatgen

import "testing"

func TestEncodeSkill(t *testing.T) {
	cases := []struct {
		given  int
		expect string
	}{
		{
			given:  743,
			expect: "[&BucCAAA=]",
		},
		{
			given:  5491,
			expect: "[&BnMVAAA=]",
		},
		{
			given:  5501,
			expect: "[&Bn0VAAA=]",
		},
	}

	for _, c := range cases {
		result := EncodeSkill(c.given)

		if result != c.expect {
			t.Errorf("EncodeSkill returned '%s' but '%s' was expected", result, c.expect)
		}
	}
}

func TestDecodeSkill(t *testing.T) {
	cases := []struct {
		given  string
		expect struct {
			id  int
			err error
		}
	}{
		{
			given: "[&BucCAAA=]",
			expect: struct {
				id  int
				err error
			}{
				id: 743,
			},
		},
		{
			given: "[&BnMVAAA=]",
			expect: struct {
				id  int
				err error
			}{
				id: 5491,
			},
		},
		{
			given: "[&Bn0VAAA=]",
			expect: struct {
				id  int
				err error
			}{
				id: 5501,
			},
		},
	}

	for _, c := range cases {
		result, err := DecodeSkill(c.given)

		if err != c.expect.err {
			t.Errorf("DecodeSkill returned an unexpected error: '%s'", err)
		}
		if result != c.expect.id {
			t.Errorf("DecodeSkill returned '%d' but '%d' was expected", result, c.expect.id)
		}
	}
}
