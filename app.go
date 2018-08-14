package wellgo

import (
	"log"
)

type App struct {
}

func Run() {
	var (
		err error
	)

	//初始化日志模块
	if err = InitLogger(); err != nil {
		log.Fatal(err)
	}
	defer CloseLogger()

	logger.Info("wellgo: initializing framework")

	//初始化配置模块
	conf := NewConfig()
	proto, err := conf.Get("config", "sys", "proto")
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	//初始化router
	InitRouter()

	switch proto {
	case "http":
		getHttpInstance().SetRPC(new(JsonRPC))
		getHttpInstance().SetProtoType(ProtoHttp)
		getHttpInstance().serveHttp()
	case "https":
		getHttpInstance().SetRPC(new(JsonRPC))
		getHttpInstance().SetProtoType(ProtoHttps)
		getHttpInstance().serveHttps()
	case "tcp":
		logger.Error("wellgo: Not support tcp now")
		panic("wellgo: Not support tcp now")
	default:
		logger.Error("wellgo: Please config your proto")
		panic("wellgo: Not support tcp now")
	}
}
