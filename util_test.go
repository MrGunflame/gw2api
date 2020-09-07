package gw2api

import "testing"

func TestConcatStrings(t *testing.T) {
	cases := []struct {
		given  []string
		expect string
	}{
		{
			given:  []string{},
			expect: "",
		},
		{
			given:  []string{"Single String"},
			expect: "Single String",
		},
		{
			given:  []string{"Hello", "World"},
			expect: "HelloWorld",
		},
	}

	for _, c := range cases {
		result := concatStrings(c.given...)

		if result != c.expect {
			t.Errorf("concatStrings returned '%s' but '%s' was expected", result, c.expect)
		}
	}
}
