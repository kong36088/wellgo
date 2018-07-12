package wellgo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

const (
	METHOD_GET   = 1
	METHOD_POST  = 2
	METHOD_PUT   = 3
	METHOD_DELTE = 4
)

var(
	appUrl string
)

func servHttp() {
	var(
		addr string
		err error
	)
	appUrl,err = conf.GetConfig("sys","app_url")
	if err != nil{
		log.Fatal(err)
	}
	addr, err = conf.GetConfig("sys", "addr")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", httpHandler)
	http.ListenAndServe(addr, nil)
}

func servHttps() {
	var (
		addr string
		cert string
		key  string
		err error
	)
	appUrl,err = conf.GetConfig("sys","app_url")
	if err != nil{
		log.Fatal(err)
	}
	addr, err = conf.GetConfig("sys", "addr")
	if err != nil {
		log.Fatal(err)
	}

	cert, err = conf.GetConfig("sys", "cert")
	if err != nil {
		log.Fatal(err)
	}
	key, err = conf.GetConfig("sys", "key")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", httpHandler)
	http.ListenAndServeTLS(addr, cert, key, nil)
}

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

func (httpReq *HttpRequest) getReqData() map[string]string {
	//TODO get data
}

type HttpResponse struct {
	Response

	header Header
}
