package wellgo

import (
	"errors"
)

var (
	OK error = errors.New("")

	ErrSystemError error = errors.New("system error")

	ErrParamError error = errors.New("param error")

	ErrValueNotFound error = errors.New("value not found")

	ErrInvalidInputFormat error = errors.New("invalid input format")

	ErrInterfaceNotFound error = errors.New("interface not found")

	ErrMap = map[error]int64{
		OK:                    0,
		ErrSystemError:        1000,
		ErrParamError:         1001,
		ErrValueNotFound:      1002,
		ErrInvalidInputFormat: 1003,
		ErrInterfaceNotFound:  1004,
	}
)

type WException struct {
	Code    int64
	Message string
}

func (we *WException) Error() string {
	return we.Message
}

func NewWException(message string, codes ...int64) WException {
	var code int64
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 0
	}

	return WException{
		Code:    code,
		Message: message,
	}
}

func RegisterError(code int64, err error) {
	ErrMap[err] = code
}

func RegisterErrorMap(em map[error]int64) {
	for k, v := range em {
		ErrMap[k] = v
	}
}

func ErrorExists(err error) bool {
	_, ex := ErrMap[err]
	return ex
}

func GetErrorCode(err error) int64 {
	code, found := ErrMap[err]
	if !found {
		return -1
	}
	return code
}

/**
 * 异常捕抓器
 */
func ErrorHandler(ctx *WContext) {
	if err := recover(); err != nil {
		var (
			code    int64
			message string
		)
		switch err.(type) {
		case *WException:
			we, _ := err.(*WException)
			code = we.Code
			message = we.Message
		case WException:
			we, _ := err.(WException)
			code = we.Code
			message = we.Message
		case error:
			e, _ := err.(error)
			code := GetErrorCode(e)
			if code == -1 {
				e = ErrSystemError
				code = GetErrorCode(e)
				message = e.Error()
			}
		default:
			logger.Error("can not handle error type", err)
			return
		}

		ctx.Proto.GetRPC().EncodeErrResponse(ctx, *NewResult(code, message))
	}
}
