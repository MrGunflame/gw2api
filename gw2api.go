package gw2api

import (
	"encoding/json"
	"net/http"
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
func (s *Session) WithEndpointAPI(endpointAPI string) *Session {
	s.endpointAPI = endpointAPI
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

// get fetches an endpoint and decodes the json response into dst
func (s *Session) get(endpoint string, dst interface{}) error {
	// Create a new request
	req, err := s.newRequest("GET", endpoint)
	if err != nil {
		return err
	}

	// Make the request
	return makeRequest(req, dst)
}

// getWithAuth fetches an endpoint that requires authentication and decodes the json response into dst
func (s *Session) getWithAuth(endpoint string, dst interface{}) error {
	// Create a new request
	req, err := s.newRequest("GET", endpoint)
	if err != nil {
		return err
	}

	// Set the AccessToken header
	if s.accessToken == "" {
		return ErrNoAccessToken
	}
	req.Header.Set("Authorization", concatStrings("Bearer ", s.accessToken))

	// Make the request
	return makeRequest(req, dst)
}

// For language endpoint
func (s *Session) getWithLang(endpoint string, dst interface{}) error {
	if strings.Contains(endpoint, "?") {
		return s.get(concatStrings(endpoint, "&lang=", s.language), &dst)
	}
	return s.get(concatStrings(endpoint, "?lang=", s.language), &dst)
}

// makeRequest takes a http request, sends it and decodes the json response into the dst interface
func makeRequest(req *http.Request, dst interface{}) error {
	// Make the http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Error checking
	var sErr *Error
	if resp.StatusCode > 299 {
		if err = json.NewDecoder(resp.Body).Decode(&sErr); err != nil {
			return err
		}
		return sErr
	}

	// Decode the json data
	if err = json.NewDecoder(resp.Body).Decode(dst); err != nil {
		return err
	}

	return nil
}

// genArgs generates the http arguments based on the given ids
func genArgs(ids ...int) string {
	switch len(ids) {
	case 0:
		return "?ids=all"
	default:
		return concatStrings("?ids=", strings.Join(itoaSlice(ids), ","))
	}
}

func genArgsString(ids ...string) string {
	switch len(ids) {
	case 0:
		return "?ids=all"
	default:
		return concatStrings("?ids=", strings.Join(ids, ","))
	}
}
