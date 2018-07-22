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

type HttpMethod uint8

var (
	http *Http
)

type HttpRequest struct {
	Request

	Url       string
	Host      string
	Uri       string
	RawInput  []byte
	Args      map[string]interface{}
	Interface string

	Method HttpMethod

	Header *Header

	ProtoType ProtoType
}

type Http struct {
	ProtoBase
}

func (httpReq *HttpRequest) GetProtoType() ProtoType {
	return httpReq.ProtoType
}

func (httpReq *HttpRequest) GetUrl() string {
	return httpReq.Url
}

func (httpReq *HttpRequest) GetHost() string {
	return httpReq.Host
}

func (httpReq *HttpRequest) GetUri() string {
	return httpReq.Uri
}

func (httpReq *HttpRequest) GetRawInput() []byte {
	return httpReq.RawInput
}

func (httpReq *HttpRequest) GetArgs() map[string]interface{} {
	return httpReq.Args
}

func (httpReq *HttpRequest) GetInterface() string {
	return httpReq.Interface
}

func (httpReq *HttpRequest) SetProtoType(protoType ProtoType) {
	httpReq.ProtoType = protoType
}

func (httpReq *HttpRequest) SetUrl(url string) {
	httpReq.Url = url
}

func (httpReq *HttpRequest) SetHost(host string) {
	httpReq.Host = host
}

func (httpReq *HttpRequest) SetUri(uri string) {
	httpReq.Uri = uri
}

func (httpReq *HttpRequest) SetRawInput(input []byte) {
	httpReq.RawInput = input
}

func (httpReq *HttpRequest) SetArgs(args map[string]interface{}) {
	httpReq.Args = args
}

func (httpReq *HttpRequest) SetInterface(interf string) {
	httpReq.Interface = interf
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
		http = &Http{}
		http.addr = addr
		http.appUrl = appUrl
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

/**
 * http 处理函数，分发请求至RPC处理器
 */
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

	if http.RPChandler == nil {
		log.Fatal("wellgo.http.RPChandler is not set")
	}

	req := &HttpRequest{
		Header: NewHeader(r.Header),
	}
	req.ProtoType = ProtoHttp
	req.Url = r.URL.String()
	req.Host = r.URL.Host
	req.Uri = r.URL.RequestURI()
	req.RawInput = b
	rsp := http.RPChandler(req)

	fmt.Println(rsp)
}

type Header struct {
	Header netHttp.Header
	//headers map[string]string
}

func NewHeader(h netHttp.Header) *Header {
	return &Header{
		Header: h,
	}
}

func (httpReq *HttpRequest) getReqData() map[string]string {
	//TODO get data
	return make(map[string]string)
}

type HttpResponse struct {
	Response

	header Header
}
