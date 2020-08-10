package gw2api

import (
	"errors"
)

// Error constants
var (
	ErrNotFound           = errors.New("not found")
	ErrInvalidAccessToken = errors.New("Invalid access token")
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
