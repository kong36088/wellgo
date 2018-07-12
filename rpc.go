package wellgo

import "fmt"

type RPC struct{

}

func (rpc *RPC) rpcHandler(input []byte) *Response{
	inputStr := string(input)
	fmt.Println(inputStr)
	return &Response{}
}