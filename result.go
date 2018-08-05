/**
 * @author wellsjiang
 */

package wellgo

import "encoding/json"

type Result struct {
	ReturnCode    int         `json:"returnCode"`
	ReturnMessage string      `json:"returnMessage"`
	Data          interface{} `json:"data"`
}

func NewResult(code int, message string, data interface{}) *Result {
	return &Result{
		ReturnCode:    code,
		ReturnMessage: message,
		Data:          data,
	}
}

func (r *Result) JsonSerialize() []byte {
	j, _ := json.Marshal(r)
	return j
}