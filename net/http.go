package net

import "github.com/kong36088/wellgo"

const (
	METHOD_GET   = 1
	METHOD_POST  = 2
	METHOD_PUT   = 3
	METHOD_DELTE = 4
)

type Header struct {
	headers map[string]string
}

func NewHeader() *Header {
	return &Header{
		headers: make(map[string]string),
	}
}
func (h *Header) Get(name string) (string, error) {
	if val, found := h.headers[name]; found {
		return val, wellgo.OK
	} else {
		return "", wellgo.ErrValueNotFound
	}
}

func (h *Header) Set(name string, value string) error {
	h.headers[name] = value
	return wellgo.OK
}

type HttpRequest struct {
	wellgo.Request

	header Header
}

type HttpResponse struct {
	wellgo.Response

	header Header
}
