/**
 * @author wellsjiang
 */

package wellgo

import "github.com/cihub/seelog"

type WContext struct {
	Proto ProtoInterface

	Req Request

	Rsp Response

	Config *Config

	Logger seelog.LoggerInterface
}

func newContext() *WContext {
	return &WContext{
	}
}

func (ctx *WContext) Write(content []byte) {
	logger.Info("write=" + string(content))

	ctx.Rsp.Write(content)
}

func (ctx *WContext) WriteString(content string) {
	logger.Info("write=" + content)

	ctx.Rsp.WriteString(content)
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

	GetRPC() RPC

	SetRPC(RPC)

	GetProtoType() ProtoType
	SetProtoType(ProtoType)
}

type Request interface {
	GetUrl() string
	GetHost() string
	GetUri() string
	GetPath() string
	GetRawInput() []byte
	GetArgs() map[string]interface{}
	GetInterface() string

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
