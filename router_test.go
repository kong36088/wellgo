/**
  * @author wellsjiang
  * @date 2018/8/14
  */

package wellgo

import "testing"

type Test struct {
	Controller
}

const (
	testStr = "test string"
)

func (t *Test) Run() *Result {
	return NewResult(123, testStr, nil)
}

func init() {
	InitRouter()
}

func TestRouter_Register(t *testing.T) {
	router.Register("a.b.c", &Test{})
}

func TestRouter_RegexpRegister(t *testing.T) {
	router.RegexpRegister("c\\.d\\..+", &Test{})
	router.RegexpRegister("open\\..+\\..+", &Test{})
}

func TestRouter_Match(t *testing.T) {
	TestRouter_Register(t)
	TestRouter_RegexpRegister(t)

	if controller, err := router.Match("a.b.c"); err != nil || controller.Run().ReturnMessage != testStr {
		t.Log(err)
		t.Log(controller)
		t.Fail()
	}

	if controller, err := router.Match("c.d.effg"); err != nil || controller.Run().ReturnMessage != testStr {
		t.Log(err)
		t.Log(controller)
		t.Fail()
	}

	if controller, err := router.Match("open.sdaf.asdjfio"); err != nil || controller.Run().ReturnMessage != testStr {
		t.Log(err)
		t.Log(controller)
		t.Fail()
	}
}
