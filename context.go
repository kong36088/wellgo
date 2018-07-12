package wellgo

import "sync"

type WContext struct {
	proto string

	req Request

	resp Response

	middlewares *sync.Map
}

func newContext(proto string, req Request, resp Response) *WContext {
	return &WContext{
		proto: proto,
		req:   req,
		resp:  resp,
		middlewares: &sync.Map{},
	}
}

func (wcont*WContext) regMiddleware(middleware *Middleware) error {
	wcont.middlewares.Store(middleware, middleware)

	return OK
}

func (wcont *WContext) delMiddleware(middleware *Middleware) error {
	wcont.middlewares.Delete(middleware)

	return OK
}

type Request struct {
	url string
	uri string
	args map[string]string

}

type Response struct {
}
