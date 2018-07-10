package wellgo

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
		middlewares & sync.Map{},
	}
}

func (wcontext *WContext) regMiddleware(middleware *Middleware) error {
	wcontext.middlewares.Store(middleware, middleware)

	return OK
}

func (wcontext *WContext) delMiddleware(middleware *Middleware) error {
	wcontext.middlewares.Del(middleware)

	return OK
}

type Request interface {
	parseRequest()
	getReqData()
}

type Response interface {
	response()
}
