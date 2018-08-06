/**
  * @author wellsjiang
  * @date 2018/8/6
  */

package wellgo

import (
	"testing"
	"github.com/bitly/go-simplejson"
	"reflect"
	"fmt"
)

type TestS struct {
	A int64 `param:"a"`

	B string `param:"b"`

	C string `param:"-"`

	D bool `param:"d",json:"test"`

	E interface{} `param:"e"`

	F interface{} `param:"f"`
}

func (t TestS) Print() {
	typ := reflect.TypeOf(t)
	v := reflect.ValueOf(t)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段
			fmt.Printf("%s %s = %v -tag:%s \n",
				typ.Field(i).Name,
				typ.Field(i).Type,
				v.Field(i).Interface(),
				typ.Field(i).Tag)
		}
	}

}

func TestAssignJsonTo(t *testing.T) {
	tt := &TestS{}
	j, _ := simplejson.NewJson([]byte(`{"a":1,"b":"b","c":"c","d":false,"e":[1,2,3,4,"mmd"],"f":{"wtf":1}}`))
	jm, _ := j.Map()

	tt.Print()
	for k, v := range jm {
		fmt.Printf("k=%s, v=%s\n", k, v)
	}
	if !AssignJsonTo(jm, reflect.ValueOf(tt), "param") {
		t.Fail()
	}
	tt.Print()
}
