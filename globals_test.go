/**
  * @author wellsjiang
  * @date 2018/8/6
  */

package wellgo

import (
	"testing"
	"reflect"
	"errors"
	"github.com/bitly/go-simplejson"
)

type TestS struct {
	A int64 `param:"a"`

	B string `param:"b"`

	C string `param:"-"`

	D bool `param:"d",json:"test"`

	E interface{} `param:"e"`

	F interface{} `param:"f"`
}

func (t TestS) Print(te *testing.T) {
	typ := reflect.TypeOf(t)
	v := reflect.ValueOf(t)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段
			te.Logf("%s %s = %v -tag:%s \n",
				typ.Field(i).Name,
				typ.Field(i).Type,
				v.Field(i).Interface(),
				typ.Field(i).Tag)
		}
	}

}

func init(){
	InitLogger()
}

func TestAssignMapTo(t *testing.T) {
	tt := &TestS{}
	j, _ := simplejson.NewJson([]byte(`{"a":1,"b":"b","c":"c","d":false,"e":[1,2,3,4,"mmd"],"f":{"mm":1}}`))
	jm, _ := j.Map()
	//jm := map[string]interface{}{"a": int8(1), "b": "1ff", "d": true,"e":"whatever","f":map[string]int{"mm":123}}

	tt.Print(t)
	for k, v := range jm {
		t.Logf("k=%s, v=%s\n", k, v)
	}
	if !AssignMapTo("bb", reflect.ValueOf(tt), "param") {
		t.Fail()
	}
	tt.Print(t)
}

func TestAssert(t *testing.T) {
	var code int64 = 123
	msg := "test exception"

	defer func() {
		if err := recover(); err != nil {
			e, _ := err.(WException)
			if e.Message != msg || e.Code != int64(code) {
				t.Log(e)
				t.Error(e.Message, e.Code)
			}
		}
	}()

	Assert(true)

	Assert(1 == 1)

	Assert(1 == 2, NewWException(errors.New(msg).Error(), code))

}
