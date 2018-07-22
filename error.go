package wellgo

import "errors"

var (
	OK error = nil

	ErrSystemError error = errors.New("system error")

	ErrParamError error = errors.New("param error")

	ErrValueNotFound error = errors.New("value not found")

	ErrInvalidInputParam error = errors.New("invalid input param")
)