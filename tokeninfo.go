package gw2api

// TokenInfo contains general token information, such as the permissions
type TokenInfo struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

// Tokeninfo returns general information about the provided token
func (s *Session) Tokeninfo() (tokeninfo TokenInfo, err error) {
	err = s.getWithAuth("/v2/tokeninfo", &tokeninfo)
	return
}
