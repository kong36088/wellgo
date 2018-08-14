/**
  * @author wellsjiang
  * @date 2018/8/14
  */

package wellgo

import (
	"testing"
)

func TestNewResult(t *testing.T) {
	NewResult(123, "456", map[string]interface{}{})
}

func TestResult_GetData(t *testing.T) {
	ResultGetDataTest(t, NewResult(123, "456", map[string]interface{}{}))
	ResultGetDataTest(t, NewResult(123, "456", &map[string]interface{}{}))
	ResultGetDataTest(t, NewResult(123, "456", "mmf"))
	ResultGetDataTest(t, NewResult(123, "456", 4))
}

func ResultGetDataTest(t *testing.T, r *Result) {
	data := r.GetData()
	data = "abc"
	t.Log(r)
	t.Log(data)
	if data == r.GetData() {
		t.Fail()
	}
}
