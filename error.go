package wellgo

import (
	"github.com/pkg/errors"
	"runtime/debug"
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

// TODO error trace back
type WException struct {
	Code int64
	Err  error
}

func (we *WException) Error() string {
	return we.Err.Error()
}

func NewWException(err error, codes ...int64) WException {
	var code int64
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 0
	}

	return WException{
		Code: code,
		Err:  err,
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
			output  []byte
		)
		logger.Error(string(debug.Stack()))
		logger.Error(err)
		switch err.(type) {
		case *WException:
			we, _ := err.(*WException)
			code = we.Code
			message = we.Error()
		case WException:
			we, _ := err.(WException)
			code = we.Code
			message = we.Error()
		case error:
			e, _ := err.(error)
			code := GetErrorCode(e)
			if code == -1 {
				e = ErrSystemError
				code = GetErrorCode(e)
				message = e.Error()
			}
		default:
			logger.Error("wellgo: can not handle error type", err)
			return
		}
		if output, err = ctx.Proto.GetRPC().EncodeErrResponse(ctx, *NewResult(code, message)); err != nil {
			logger.Error(err)
		}
		ctx.Write(output)
	}
}
