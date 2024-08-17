package errs

import "errors"

var (
	ErrUnsupportedTokenType = errors.New("token type is not supported")
	ErrTokenNotFound        = errors.New("token not found")
	ErrInvalidUUIDFormat    = errors.New("invalid UUID format")
)
