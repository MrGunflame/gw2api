package gw2api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var defaultEndpointAPI = "https://api.guildwars2.com"

// Session is used to make requests to the gw2api. It also handles authentication.
type Session struct {
	endpointAPI string
	accessToken string
	language    string
}

// New creates a new gw2api session
func New() *Session {
	return &Session{
		endpointAPI: defaultEndpointAPI,
		language:    "en",
	}
}

// WithEndpointAPI sets the API endpoint
func (s *Session) WithEndpointAPI(endpointApi string) *Session {
	s.endpointAPI = endpointApi
	return s
}

// WithAccessToken sets an accesstoken for a session
func (s *Session) WithAccessToken(token string) *Session {
	s.accessToken = token
	return s
}

// WithLanguage sets the sessions language
func (s *Session) WithLanguage(lang string) *Session {
	s.language = lang
	return s
}

// newRequest creates a new request with the json content type and the valid access token
func (s *Session) newRequest(method, endpoint string) (*http.Request, error) {
	req, err := http.NewRequest(method, concatStrings(s.endpointAPI, endpoint), nil)
	if err != nil {
		return nil, err
	}

	// Add the accesstoken to the authentication header
	if s.accessToken != "" {
		req.Header.Set("Authorization", concatStrings("Bearer ", s.accessToken))
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

	return nil
}

func (s *Session) getWithAuth(endpoint string, dst interface{}) error {
	return s.get(endpoint, dst)
}

func decode(src io.Reader, dst interface{ Err() error }) error {
	if err := json.NewDecoder(src).Decode(dst); err != nil {
		return err
	}
	return dst.Err()
}

// genArgs generates the http arguments based on the given ids
func genArgs(ids ...int) string {
	switch len(ids) {
	case 0:
		return "?ids=all"
	case 1:
		return concatStrings("?id=", strconv.Itoa(ids[0]))
	default:
		return concatStrings("?ids=", strings.Join(itoaSlice(ids), ","))
	}
}
