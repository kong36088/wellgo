package wellgo

import (
	"fmt"
	"io/ioutil"
	netHttp "net/http"
	"log"
)

const (
	METHOD_GET   = 1
	METHOD_POST  = 2
	METHOD_PUT   = 3
	METHOD_DELTE = 4
)

var (
	http *Http
)

type Http struct {
	addr string

	appUrl string
}

func getHttpInstance() *Http {
	if http == nil {
		var (
			appUrl string
			addr   string
			err    error
		)
		appUrl, err = conf.GetConfig("sys", "app_url")
		if err != nil {
			log.Fatal(err)
		}
		addr, err = conf.GetConfig("sys", "addr")
		if err != nil {
			log.Fatal(err)
		}
		http = &Http{
			addr:   addr,
			appUrl: appUrl,
		}
	}
	return http
}

func (http *Http) serveHttp() {
	netHttp.HandleFunc("/", http.httpHandler)
	netHttp.ListenAndServe(http.addr, nil)
}

func (http *Http) serveHttps() {
	var (
		cert string
		key  string
		err  error
	)
	cert, err = conf.GetConfig("sys", "cert")
	if err != nil {
		log.Fatal(err)
	}
	key, err = conf.GetConfig("sys", "key")
	if err != nil {
		log.Fatal(err)
	}

	netHttp.HandleFunc("/", http.httpHandler)
	netHttp.ListenAndServeTLS(http.addr, cert, key, nil)
}

func (http *Http) httpHandler(w netHttp.ResponseWriter, r *netHttp.Request) {
	if r.RequestURI != http.appUrl {
		netHttp.NotFound(w, r)
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
