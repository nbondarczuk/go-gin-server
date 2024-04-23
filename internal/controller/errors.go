package controller

import "errors"

var (
	ErrBackendClientConnect = errors.New("controller backend client connection error")
	ErrBackendClientPing    = errors.New("controller backend client ping error")
)
