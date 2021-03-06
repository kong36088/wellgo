/**
 * @author wellsjiang
 */

package wellgo

import (
	"io/ioutil"
	netHttp "net/http"
	"log"
	"reflect"
	"github.com/kong36088/wellgo/utils"
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

	rpc RPC

	conf *Config

	ProtoType ProtoType
}

func (http *Http) Addr() string {
	return http.addr
}

func (http *Http) AppUrl() string {
	return http.appUrl
}

func (http *Http) GetRPC() RPC {
	return http.rpc
}

func (http *Http) SetRPC(rpc RPC) {
	http.rpc = rpc
}
func (http *Http) GetProtoType() ProtoType {
	return http.ProtoType
}
func (http *Http) SetProtoType(protoType ProtoType) {
	http.ProtoType = protoType
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

	R *netHttp.Request
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

	W netHttp.ResponseWriter
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
func (rsp *HttpResponse) Write(content []byte) {
	_, err := rsp.W.Write(content)
	if err != nil {
		logger.Error(err)
	}
}
func (rsp *HttpResponse) WriteString(content string) {
	_, err := rsp.W.Write([]byte(content))
	if err != nil {
		logger.Error(err)
	}
}

func getHttpInstance() *Http {
	if http == nil {
		var (
			appUrl string
			addr   string
			err    error
		)
		http = &Http{
			conf: NewConfig(),
		}
		appUrl, err = http.conf.Get("config", "sys", "app_url")
		if err != nil {
			logger.Critical(err)
			panic(err)
		}
		addr, err = http.conf.Get("config", "sys", "addr")
		if err != nil {
			logger.Critical(err)
			panic(err)
		}

		http.addr = addr
		http.appUrl = appUrl
	}
	return http
}

func (http *Http) serveHttp() {
	logger.Infof("wellgo: start listening %s, uri=%s, proto=http", http.addr, http.appUrl)
	netHttp.HandleFunc("/", http.httpHandler)
	netHttp.ListenAndServe(http.addr, nil)
}

func (http *Http) serveHttps() {
	var (
		cert string
		key  string
		err  error
	)
	if cert, err = http.conf.Get("config", "sys", "cert"); err != nil {
		log.Fatal(err)
	}

	if key, err = http.conf.Get("config", "sys", "key"); err != nil {
		log.Fatal(err)
	}

	logger.Infof("wellgo: start listening %s, uri=%s, proto=https", http.addr, http.appUrl)
	netHttp.HandleFunc("/", http.httpHandler)
	netHttp.ListenAndServeTLS(http.addr, cert, key, nil)
}

//TODO write headers
/**
 * http 处理函数
 */
func (http *Http) httpHandler(w netHttp.ResponseWriter, r *netHttp.Request) {
	timer = &utils.Timer{}
	timer.Start()

	// init Req
	req := &HttpRequest{
		Header: NewHeader(r.Header),
		Url:    r.URL.String(),
		Uri:    r.URL.RequestURI(),
		Host:   r.URL.Host,
		R:      r,
	}
	// init rsp
	rsp := &HttpResponse{
		W: w,
	}
	// init ctx
	ctx := newContext()
	ctx.Config = NewConfig()
	ctx.Proto = http
	ctx.Req = req
	ctx.Rsp = rsp
	ctx.Logger = GetLoggerInstance()

	// error handler
	defer ErrorHandler(ctx)

	// judge url
	if r.RequestURI != http.appUrl {
		netHttp.NotFound(w, r)
		return
	}
	// read request body
	b, err := ioutil.ReadAll(r.Body)
	Assert(err == nil, NewWException(err))

	logger.Infof("wellgo: start, req=%s", b)

	if http.rpc == nil {
		logger.Critical("wellgo: wellgo.http.rpc is not set")
		panic("wellgo.http.rpc is not set")
	}

	req.RawInput = b

	// process rpc
	parsedReq, err := http.rpc.RPCHandler(req)
	Assert(err == nil, NewWException(err))

	req = parsedReq.(*HttpRequest)

	// router
	controller, err := router.Match(req.GetInterface())
	Assert(err == nil, NewWException(err))

	// controller process
	controller.Init(ctx)

	AssignMapTo(ctx.Req.GetArgs(), reflect.ValueOf(controller), "param")

	result := controller.Run()

	output, err := http.GetRPC().EncodeResponse(ctx, *result)
	if err != nil {
		logger.Error(err)
	}
	ctx.Write(output)
}

type HttpHeader struct {
	*netHttp.Header
}

func NewHeader(h netHttp.Header) *HttpHeader {
	return &HttpHeader{}
}
