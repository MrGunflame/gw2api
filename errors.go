package gw2api

import (
	"errors"
)

// Error constants
var (
	ErrNotFound           = errors.New("not found")
	ErrNoAccessToken      = errors.New("No access token provided")
	ErrInvalidAccessToken = errors.New("Invalid access token")
	ErrCannotUseAll       = errors.New("unable to use 'all' keyword for this API")
)

// Error is an error returned from the gw2api
type Error struct {
	Err  string `json:"error"`
	Text string `json:"text"`
}

func (e *Error) Error() string {
	if e.Err != "" {
		return e.Err
	}
	return e.Text
}
