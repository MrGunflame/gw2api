package gw2api

type TokenInfo struct {
  Error
  ID string `json:"id"`
  Name string `json:"name"`
  Permissions []string `json:"permissions"`
}

func (s *Session) GetTokenInfo() (tokeninfo TokenInfo, err error) {

  err = s.get("/v2/tokeninfo", &tokeninfo)
  if err != nil {
    return
  }

  return
}
