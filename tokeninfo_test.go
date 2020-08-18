package gw2api

import "testing"

func TestTokenInfo(t *testing.T) {
	apiAuth := New().WithAccessToken(testAccessToken)

	if _, err := apiAuth.GetTokenInfo(); err != nil {
		return
	}
}
