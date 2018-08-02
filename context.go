/**
 * @author wellsjiang
 */

package wellgo

import "sync"

type WContext struct {
	proto ProtoInterface

	req Request

	resp Response

	middlewares *sync.Map
}

func newContext(proto ProtoInterface, req Request, resp Response) *WContext {
	return &WContext{
		proto: proto,
		req:   req,
		resp:  resp,
	}
}

const (
	ProtoHttp  = 1
	ProtoHttps = 2
	ProtoTcp   = 3
)

type ProtoType uint8

type ProtoInterface interface {
	Addr() string

	AppUrl() string

	RPCHandler() (func(Request) (Request, error))

	SetRPCHandler(func(Request) (Request, error))
}

type Request interface {
	GetProtoType() ProtoType
	GetUrl() string
	GetHost() string
	GetUri() string
	GetPath() string
	GetRawInput() []byte
	GetArgs() map[string]interface{}
	GetInterface() string

	SetProtoType(ProtoType)
	SetUrl(string)
	SetHost(string)
	SetUri(string)
	SetPath(string)
	SetRawInput([]byte)
	SetArgs(map[string]interface{})
	SetInterface(string)
}

type Response interface {
	GetReturnCode() int
	GetReturnMessage() string
	GetData() interface{}

	SetReturnCode(int)
	SetReturnMessage(string)
	SetData(interface{})
}
