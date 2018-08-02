/**
 * @author wellsjiang
 */

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

type Http struct {
	addr string

	appUrl string

	rpcHandler func(Request) (Request, error)
}

type HttpRequest struct {
	Url       string
	Host      string
	Uri       string
	Path      string
	RawInput  []byte
	Args      map[string]interface{}
	Interface string

	Method HttpMethod

	Header *HttpHeader

	ProtoType ProtoType
}

func (http *Http) Addr() string {
	return http.addr
}

func (http *Http) AppUrl() string {
	return http.appUrl
}

func (http *Http) RPCHandler() func(Request) (Request, error) {
	return http.rpcHandler
}

func (http *Http) SetRPCHandler(rpcHandler func(Request) (Request, error)) {
	http.rpcHandler = rpcHandler
}

func (req *HttpRequest) GetProtoType() ProtoType {
	return req.ProtoType
}

func (req *HttpRequest) GetUrl() string {
	return req.Url
}

func (req *HttpRequest) GetHost() string {
	return req.Host
}

func (req *HttpRequest) GetUri() string {
	return req.Uri
}

func (req *HttpRequest) GetPath() string {
	return req.Path
}

func (req *HttpRequest) GetRawInput() []byte {
	return req.RawInput
}

func (req *HttpRequest) GetArgs() map[string]interface{} {
	return req.Args
}

func (req *HttpRequest) GetInterface() string {
	return req.Interface
}

func (req *HttpRequest) GetHeader() HttpHeader {
	return req.Header
}

func (req *HttpRequest) SetProtoType(protoType ProtoType) {
	req.ProtoType = protoType
}

func (req *HttpRequest) SetUrl(url string) {
	req.Url = url
}

func (req *HttpRequest) SetHost(host string) {
	req.Host = host
}

func (req *HttpRequest) SetUri(uri string) {
	req.Uri = uri
}

func (req *HttpRequest) SetPath(path string) {
	req.Path = path
}

func (req *HttpRequest) SetRawInput(input []byte) {
	req.RawInput = input
}

func (req *HttpRequest) SetArgs(args map[string]interface{}) {
	req.Args = args
}

func (req *HttpRequest) SetInterface(interf string) {
	req.Interface = interf
}

type HttpResponse struct {
	ReturnCode    int
	ReturnMessage string
	Data          interface{}
	Header        *HttpHeader
}

func (rsp *HttpResponse) GetReturnCode() int {
	return rsp.ReturnCode
}
func (rsp *HttpResponse) GetReturnMessage() string {
	return rsp.ReturnMessage
}
func (rsp *HttpResponse) GetData() interface{} {
	return rsp.Data
}
func (rsp *HttpResponse) GetHeader() *HttpHeader {
	return rsp.Header
}
func (rsp *HttpResponse) SetReturnCode(code int) {
	rsp.ReturnCode = code
}
func (rsp *HttpResponse) SetReturnMessage(message string) {
	rsp.ReturnMessage = message
}
func (rsp *HttpResponse) SetData(data interface{}) {
	rsp.Data = data
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
 * http 处理函数
 */
func (http *Http) httpHandler(w netHttp.ResponseWriter, r *netHttp.Request) {
	var (
		parsedReq  Request
		controller *Controller
		ctx        *WContext
	)

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

	if http.rpcHandler == nil {
		log.Fatal("wellgo.http.rpcHandler is not set")
	}

	// init req
	req := &HttpRequest{
		Header: NewHeader(r.Header),
	}
	req.ProtoType = ProtoHttp
	req.Url = r.URL.String()
	req.Host = r.URL.Host
	req.Uri = r.URL.RequestURI()
	req.RawInput = b

	parsedReq, err = http.rpcHandler(req)
	if err != nil {
		// TODO error handler
		return
	}

	req = parsedReq.(*HttpRequest)

	controller, err = router.Match(req.GetPath())
	if err != nil {
		return
	}

	//init rsp
	rsp := &HttpResponse{}

	ctx = newContext(http, req, rsp)

	controller.Init(ctx)

	controller.Run()

}

type HttpHeader struct {
	*netHttp.Header
}

func NewHeader(h netHttp.Header) *HttpHeader {
	return &HttpHeader{
		Header: h,
	}
}
