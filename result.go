/**
 * @author wellsjiang
 */

package wellgo

import (
	"encoding/json"
)

type Result struct {
	ReturnCode    int64       `json:"returnCode"`
	ReturnMessage string      `json:"returnMessage"`
	Data          interface{} `json:"data"`
}

func NewResult(code int64, message string, data ...interface{}) *Result {
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	} else {
		d = nil
	}

	return &Result{
		ReturnCode:    code,
		ReturnMessage: message,
		Data:          d,
	}
}

func (r *Result) GetData() interface{} {
	// return copy value
	return r.Data
}

func (r *Result) GetCode() int64 {
	return r.ReturnCode
}

func (r *Result) GetMessage() string {
	return r.ReturnMessage
}

func (r *Result) JsonSerialize() []byte {
	j, _ := json.Marshal(r)
	return j
}
