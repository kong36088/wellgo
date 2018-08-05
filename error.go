package wellgo

import (
	"errors"
)

var (
	OK error = nil

	ErrSystemError error = errors.New("system error")

	ErrParamError error = errors.New("param error")

	ErrValueNotFound error = errors.New("value not found")

	ErrInvalidInputParam error = errors.New("invalid input param")
)

/**
 * 异常捕抓器
 */
func ErrorHandler(req Request, rsp Response) {
	if err := recover(); err != nil {
		rsp.WriteString(err.(string))
	}
}
