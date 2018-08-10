/**
  * @author wellsjiang
  * @date 2018/8/11
  */

package wellgo

import (
	"testing"
	"errors"
)

func TestErrorExists(t *testing.T) {
	if !ErrorExists(OK) {
		t.Fail()
	}
	if ErrorExists(errors.New("some error here")) {
		t.Fail()
	}
}

func TestGetErrorCode(t *testing.T) {
	if GetErrorCode(OK) != 0 {
		t.Fail()
	}
	if GetErrorCode(ErrSystemError) != 1000 {
		t.Fail()
	}
	for e, c := range ErrMap {
		if GetErrorCode(e) != c {
			t.Fail()
		}
	}
	if GetErrorCode(errors.New("some error here")) != -1 {
		t.Fail()
	}
}

func TestRegisterError(t *testing.T) {
	e := errors.New("some test error")
	RegisterError(123456, e)
	if !ErrorExists(e) {
		t.Fail()
	}
	if GetErrorCode(e) != 123456 {
		t.Fail()
	}
}

func TestRegisterErrorMap(t *testing.T) {
	ea, eb := errors.New("err test1"), errors.New("err test2")
	em := map[error]int64{
		ea: 1234567,
		eb: 1234568,
	}
	RegisterErrorMap(em)

	if !ErrorExists(ea) || !ErrorExists(eb) {
		t.Fail()
	}
	if GetErrorCode(ea) != em[ea] || GetErrorCode(eb) != em[eb] {
		t.Fail()
	}
}
