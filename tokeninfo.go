package gw2api

// TokenInfo contains general token information, such as the permissons
type TokenInfo struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Permissions []string `json:"permissions"`
}

// GetTokenInfo returns general information about the provided token
func (s *Session) GetTokenInfo() (tokeninfo TokenInfo, err error) {
  err = s.get("/v2/tokeninfo", &tokeninfo)
  return
}
