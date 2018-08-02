/**
 * @author wellsjiang
 */

package wellgo

import (
	"encoding/json"
)

type RPC interface {
	RPCHandler(Request) (Request, error)

	EncodeResponse(Request, Response) ([]byte, error)

	EncodeErrResponse(Request, Response, error) ([]byte, error)
}

type JsonRPCReq struct {
	Id      int64       `json:"id"`
	Version float32     `json:"jsonrpc"`
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
	var (
		input    JsonRPCReq
		inputMap map[string]interface{}
	)

	//读取json入参
	err := json.Unmarshal(req.GetRawInput(), input)
	if err != nil {
		return nil, err
	}
	req.SetInterface(input.Method)

	//处理入参
	args := make(map[string]interface{})
	inputMap, err = input.Param.(map[string]interface{})
	if err != nil {
		return nil, ErrInvalidInputParam
	}
	for k, v := range inputMap {
		args[k] = v
	}
	req.SetArgs(args)

	return req, OK
}

func (j *JsonRPC) EncodeResponse(req Request, rsp Response) ([]byte, error) {
	var input JsonRPCReq
	err := json.Unmarshal(req.GetRawInput(), input)
	if err != nil {
		logger.Error(err)
	}
	return json.Marshal(JsonRPCRsp{
		Id:     input.Id,
		Error:  nil,
		Result: rsp.GetData(),
	})
}

func (j *JsonRPC) EncodeErrResponse(req Request, rsp Response, err error) ([]byte, error) {
	var input JsonRPCReq
	err = json.Unmarshal(req.GetRawInput(), input)
	if err != nil {
		logger.Error(err)
		return []byte(""), err
	}
	return json.Marshal(JsonRPCRsp{
		Id:     input.Id,
		Error:  rsp.GetReturnMessage(),
		Result: rsp.GetData(),
	})
}
