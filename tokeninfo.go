package gw2api

type TokenInfo struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Permissions []string `json:"permissions"`
}

func (s *Session) GetTokenInfo() (tokeninfo TokenInfo, err error) {
  err = s.get("/v2/tokeninfo", &tokeninfo)
  return
}
