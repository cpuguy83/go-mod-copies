package platforms

import "errors"

var (
	ErrInvalid        = errors.New("invalid argument")
	ErrNotFound       = errors.New("not found")
	ErrNotImplemented = errors.New("not implemented")
)
