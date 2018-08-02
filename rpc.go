/**
 * @author wellsjiang
 */

package wellgo

import (
	"fmt"
	"encoding/json"
)

var (
	rpc *RPC
)

type RPC struct {
}

func getRPCInstance() *RPC {
	if rpc == nil {
		rpc = &RPC{}
	}
	return rpc
}

func (rpc *RPC) rpcHandler(req Request) Request {
	inputStr := string(req.GetRawInput())
	fmt.Println(inputStr)
	return req
}

type JsonRpcReq struct {
	Id      int64       `json:"id"`
	Version float32     `json:"jsonrpc"`
	Method  string      `json:"method"`
	Param   interface{} `json:"param"`
}

type JsonRpcRsp struct {
	Id     int64       `json:"id"`
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
}

// TODO 规范JSON-RPC返回格式
func (rpc *RPC) jsonRPCHandler(req Request) (Request, error) {
	var (
		input    JsonRpcReq
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

func (rpc *RPC) jsonRPCResponseHandler(req Request, rsp Response) (string, error) {

}
