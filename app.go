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
	if err = GetLoggerInstance().Init(); err != nil {
		log.Fatal(err)
	}
	defer logger.Close()

	//初始化配置模块
	conf := NewConfig()
	proto, err := conf.Get("config", "sys", "Proto")
	if err != nil {
		log.Fatal(err)
	}

	switch proto {
	case "http":
		getHttpInstance().SetRPC(new(JsonRPC))
		getHttpInstance().serveHttp()
	case "https":
		getHttpInstance().SetRPC(new(JsonRPC))
		getHttpInstance().serveHttps()
	case "tcp":
	default:
		log.Fatal("Please config your Proto")
	}
}
