/**
  * @author wellsjiang
  * @date 2018/8/3
  */

package wellgo

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
)

type JsonRPCReq struct {
	Id      string      `json:"id"`
	Version float64     `json:"jsonrpc"`
	Method  string      `json:"method"`
	Param   interface{} `json:"param"`
}

type JsonRPCRsp struct {
	Id      string        `json:"id"`
	Version float64       `json:"jsonrpc"`
	Error   JsonRPCRspErr `json:"error"`
	Result  interface{}   `json:"result"`
}

type JsonRPCRspErr struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type JsonRPC struct{}

func (j *JsonRPC) RPCHandler(req Request) (Request, error) {
	//读取json入参
	input, err := simplejson.NewJson(req.GetRawInput())
	if err != nil {
		return nil, err
	}

	//判断Id
	if input.Get("id").MustInt(0) == 0 {
		return nil, ErrInvalidInputFormat
	}
	// 判断version
	if input.Get("jsonrpc").MustFloat64(0) != 2.0 {
		return nil, ErrInvalidInputFormat
	}

	//处理入参
	req.SetInterface(input.Get("method").MustString(""))

	args := input.Get("param").MustMap(make(map[string]interface{}))
	if err != nil {
		return nil, ErrInvalidInputFormat
	}

	req.SetArgs(args)

	return req, OK
}

//TODO encode result
func (j *JsonRPC) EncodeResponse(ctx *WContext, result Result) ([]byte, error) {
	var output []byte
	input, err := simplejson.NewJson(ctx.Req.GetRawInput())
	if err != nil {
		output, _ = json.Marshal(JsonRPCRsp{
			Id:      input.Get("id").MustString(),
			Version: 2.0,
			Result:  ctx.Rsp.GetData(),
		})
		logger.Error(err)
		return output, err
	}

	return json.Marshal(JsonRPCRsp{
		Id:      input.Get("id").MustString(),
		Version: 2.0,
		Result:  ctx.Rsp.GetData(),
	})
}

func (j *JsonRPC) EncodeErrResponse(ctx *WContext, result Result) ([]byte, error) {
	var id string
	input, err := simplejson.NewJson(ctx.Req.GetRawInput())
	if err == nil {
		id = input.Get("id").MustString()
	}

	return json.Marshal(JsonRPCRsp{
		Id:      id,
		Version: 2.0,
		Error: JsonRPCRspErr{
			Code:    result.GetCode(),
			Message: result.GetMessage(),
		},
	})
}
