package gw2api

import (
  "net/http"
  "encoding/json"
  "web/pkg/util"
  "io"
  "strconv"
  "strings"
)

var defaultEndpointAPI = "https://api.guildwars2.com"

type Session struct {
  EndpointAPI string
  AccessToken string
  Language string
}

func New() *Session {
  return &Session{
    EndpointAPI: defaultEndpointAPI,
    Language: "en",
  }
}

func (s *Session) WithEndpointAPI(endpointApi string) *Session {
  s.EndpointAPI = endpointApi
  return s
}

func (s *Session) WithAccessToken(token string) *Session {
  s.AccessToken = token
  return s
}

func (s *Session) WithLanguage(lang string) *Session {
  s.Language = lang
  return s 
}

// newRequest creates a new request with the json content type and the valid access token
func (s *Session) newRequest(method, endpoint string) (*http.Request, error) {
  req, err := http.NewRequest(method, util.ConcatStrings(s.EndpointAPI, endpoint), nil)
  if err != nil {
    return nil, err
  }

  // Add the accesstoken to the authentication header
  if s.AccessToken != "" {
    req.Header.Set("Authorization", util.ConcatStrings("Bearer ", s.AccessToken))
  }

  return req, nil
}

func (s *Session) get(endpoint string, dst interface{}) error {

  req, err := s.newRequest("GET", endpoint)
  if err != nil {
    return err
  }

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  if err = json.NewDecoder(resp.Body).Decode(dst); err != nil {
    return err
  }

  return dst.Err()
}

func (s *Session) getWithAuth(endpoint string, dst interface{}) error {
  return s.get(endpoint string, dst interface{})
}

func decode(src io.Reader, dst interface{Err() error}) error {
  if err := json.NewDecoder(src).Decode(dst); err != nil {
    return err
  }
  return dst.Err()
}

// genArgs generates the http arguments based on the given ids
func genArgs(ids...int) string {
  switch len(ids) {
  case 0:
    return "?ids=all"
  case 1:
    return util.ConcatStrings("?id=", strconv.Itoa(ids[0]))
  default:
    return util.ConcatStrings("?ids=", strings.Join(util.ItoaSlice(ids), ","))
  }
}

func makeURL(args...string) string {
  return util.ConcatStrings(args...)
}
