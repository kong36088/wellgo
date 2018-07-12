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
	if err = InitConfig(); err != nil {
		log.Fatal(err)
	}
	proto, err := conf.GetConfig("sys", "proto")
	if err != nil {
		log.Fatal(err)
	}

	switch proto {
	case "http":
		getHttpInstance().serveHttp()
	case "https":
		getHttpInstance().serveHttps()
	case "tcp":
	default:
		log.Fatal("Please config your proto")
	}
}
