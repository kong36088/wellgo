package wellgo

import (
	"errors"
)

var (
	OK error = nil

	ErrSystemError error = errors.New("system error")

	ErrParamError error = errors.New("param error")

	ErrValueNotFound error = errors.New("value not found")

	ErrInvalidInputFormat error = errors.New("invalid input format")

	ErrInterfaceNotFound error = errors.New("interface not found")

	ErrMap = map[error]int{
		OK:                    0,
		ErrSystemError:        1000,
		ErrParamError:         1001,
		ErrValueNotFound:      1002,
		ErrInvalidInputFormat: 1003,
		ErrInterfaceNotFound:  1004,
	}
)

func RegisterError(code int, err error) {
	ErrMap[err] = code
}

func RegisterErrorMap(em map[error]int) {
	for k, v := range em {
		ErrMap[k] = v
	}
}

/**
 * 异常捕抓器
 */
func ErrorHandler(req Request, rsp Response) {
	if err := recover(); err != nil {
		rsp.WriteString(err.(string))
	}
}
