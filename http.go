package wellgo

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

const (
	METHOD_GET   = 1
	METHOD_POST  = 2
	METHOD_PUT   = 3
	METHOD_DELTE = 4
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != appUrl {
		http.NotFound(w, r)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	//TODO ASSERT TYPE?
	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Printf("%s", b)
}

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
		return val, OK
	} else {
		return "", ErrValueNotFound
	}
}

func (h *Header) Set(name string, value string) error {
	h.headers[name] = value
	return OK
}

type HttpRequest struct {
	Request

	header Header
}

type HttpResponse struct {
	Response

	header Header
}
