package gw2api

import "testing"

func TestNew(t *testing.T) {
	api := New()

	endpoint := "https://guilwars2.com"
	api = api.WithEndpointAPI(endpoint)
	if api.endpointAPI != endpoint {
		t.Error("WithNewEndpointAPI failed")
	}

	token := "test"
	api = api.WithAccessToken(token)
	if api.accessToken != token {
		t.Error("WithNewAccessToken failed")
	}

	lang := "de"
	api = api.WithLanguage(lang)
	if api.language != lang {
		t.Error("WithLanguage failed")
	}
}

func TestGenArgs(t *testing.T) {
	cases := []struct {
		given  []int
		expect string
	}{
		{
			given:  []int{},
			expect: "?ids=all",
		},
		{
			given:  []int{1, 2},
			expect: "?ids=1,2",
		},
	}

	for _, c := range cases {
		result := genArgs(c.given...)
		if result != c.expect {
			t.Error("genArgs failed: EXPECT: ", c.expect, " GOT: ", result)
		}
	}
}

func TestGenArgsString(t *testing.T) {
	cases := []struct {
		given  []string
		expect string
	}{
		{
			given:  []string{},
			expect: "?ids=all",
		},
		{
			given:  []string{"A", "B"},
			expect: "?ids=A,B",
		},
	}

	for _, c := range cases {
		result := genArgsString(c.given...)
		if result != c.expect {
			t.Error("genArgsString failed: EXPECT: ", c.expect, " GOT: ", result)
		}
	}
}
