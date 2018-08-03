/**
 * @author wellsjiang
 */

package wellgo

type WContext struct {
	Proto ProtoInterface

	Req Request

	Rsp Response
}

func newContext(proto ProtoInterface, req Request, resp Response) *WContext {
	return &WContext{
		Proto: proto,
		Req:   req,
		Rsp:   resp,
	}
}

const (
	ProtoHttp  = 1
	ProtoHttps = 2
	ProtoTcp   = 3
)

type ProtoType uint8

type ProtoInterface interface {
	Addr() string

	AppUrl() string

	RPC() RPC

	SetRPC(RPC)
}

type Request interface {
	GetProtoType() ProtoType
	GetUrl() string
	GetHost() string
	GetUri() string
	GetPath() string
	GetRawInput() []byte
	GetArgs() map[string]interface{}
	GetInterface() string

	SetProtoType(ProtoType)
	SetUrl(string)
	SetHost(string)
	SetUri(string)
	SetPath(string)
	SetRawInput([]byte)
	SetArgs(map[string]interface{})
	SetInterface(string)
}

type Response interface {
	GetReturnCode() int
	GetReturnMessage() string
	GetData() interface{}

	SetReturnCode(int)
	SetReturnMessage(string)
	SetData(interface{})

	Write([]byte)
	WriteString(string)
}
