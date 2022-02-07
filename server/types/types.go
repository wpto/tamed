package types

import "errors"

var (
	ErrNotFound   = errors.New("resource not found")
	ErrBadRequest = errors.New("bad request")
)
