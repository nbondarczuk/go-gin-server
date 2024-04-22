package controller

import "errors"

var (
	ErrClientConnect = errors.New("client connection error")
	ErrClientPing    = errors.New("client ping error")
)
