package bitly

import "errors"

var (
	ErrEncodingError = errors.New("unable to encode to buffer")
	ErrRequestError  = errors.New("unable to create request")
)
