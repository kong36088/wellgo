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
	Id      int64       `json:"id"`
	Version float64     `json:"jsonrpc"`
	Method  string      `json:"method"`
	Param   interface{} `json:"param"`
}

type JsonRPCRsp struct {
	Id     int64       `json:"id"`
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
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

func (j *JsonRPC) EncodeResponse(req Request, rsp Response) ([]byte, error) {
	input, err := simplejson.NewJson(req.GetRawInput())
	if err != nil {
		logger.Error(err)
		return []byte(""), err
	}
	return json.Marshal(JsonRPCRsp{
		Id:     input.Get("id").MustInt64(0),
		Error:  nil,
		Result: rsp.GetData(),
	})
}

func (j *JsonRPC) EncodeErrResponse(req Request, rsp Response, err error) ([]byte, error) {
	input, err := simplejson.NewJson(req.GetRawInput())
	if err != nil {
		logger.Error(err)
		return []byte(""), err
	}
	return json.Marshal(JsonRPCRsp{
		Id:     input.Get("id").MustInt64(0),
		Error:  rsp.GetReturnMessage(),
		Result: rsp.GetData(),
	})
}
