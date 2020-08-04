package gw2api

import (
  "errors"
)

var (
  ErrNotFound = errors.New("not found")
  ErrInvalidAccessToken = errors.New("Invalid access token")
)

type Error struct {
  HTTPError string `json:"error"`
  Text string `json:"text"`
}

func (e *Error) Err() error {
  switch {
  case e.HTTPError == "not found":
    return ErrNotFound
  case e.Text == "Invalid access token":
    return ErrInvalidAccessToken
  }
  return nil
}
