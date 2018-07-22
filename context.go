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

type ProtoBase struct {
	addr string

	appUrl string

	RPChandler func(*Request) *Response
}

func (proto *ProtoBase) SetRPCHandler(rpcHandler func(*Request) *Response) {
	proto.RPChandler = rpcHandler
}

type Request struct {
	Url      string
	Host     string
	Uri      string
	RawInput []byte
	Args     map[string]interface{}
	Interface string
}

type Response struct {
	returnCode    int
	returnMessage string
	data          interface{}
}
