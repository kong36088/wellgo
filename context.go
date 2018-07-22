package wellgo

import "sync"

type WContext struct {
	proto *ProtoBase

	req Request

	resp Response

	middlewares *sync.Map
}

func newContext(proto *ProtoBase, req Request, resp Response) *WContext {
	return &WContext{
		proto:       proto,
		req:         req,
		resp:        resp,
		middlewares: &sync.Map{},
	}
}

func (wcont *WContext) regMiddleware(middleware *Middleware) error {
	wcont.middlewares.Store(middleware, middleware)

	return OK
}

func (wcont *WContext) delMiddleware(middleware *Middleware) error {
	wcont.middlewares.Delete(middleware)

	return OK
}

const (
	ProtoHttp  = 1
	ProtoHttps = 2
	ProtoTcp   = 3
)

type ProtoType uint8

type ProtoBase struct {
	addr string

	appUrl string

	RPChandler func(Request) *Response
}

func (proto *ProtoBase) SetRPCHandler(rpcHandler func(Request) *Response) {
	proto.RPChandler = rpcHandler
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

type Response struct {
	returnCode    int
	returnMessage string
	data          interface{}
}
