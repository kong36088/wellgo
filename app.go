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
		log.Fatal("wellgo: Not support tcp now")
	default:
		log.Fatal("wellgo: Please config your proto")
	}
}
