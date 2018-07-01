package wellgo

type WContext struct {
	proto string

	req Request

	resp Response
}

func newContext(proto string, req Request, resp Response) *WContext{
	return &WContext{
		proto: proto,
		req:   req,
		resp:  resp,
	}
}

type Request interface {
	parseRequest()
	getArgs()
}

type Response interface {
	response()
}
