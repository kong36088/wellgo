package wellgo

import (
	"log"
	"net/http"
)

var (
	appUrl string
	addr   string
)

type App struct {
}

func Run() {
	var (
		err error
	)

	//初始化日志模块
	if err = GetLogger().Init(); err != nil {
		log.Fatal(err)
	}
	defer logger.Close()

	//初始化配置模块
	if err = InitConfig(); err != nil {
		log.Fatal(err)
	}

	appUrl, err = conf.GetConfig("sys", "app_url")
	if err != nil {
		log.Fatal(err)
	}
	addr, err = conf.GetConfig("sys", "addr")
	if err != nil {
		log.Fatal(err)
	}
	proto, err := conf.GetConfig("sys", "proto")
	if err != nil {
		log.Fatal(err)
	}

	switch proto {
	case "http":
		http.HandleFunc("/", httpHandler)
		http.ListenAndServe(addr, nil)
	case "https":
		var (
			cert string
			key  string
		)
		cert, err = conf.GetConfig("sys", "cert")
		if err != nil {
			log.Fatal(err)
		}
		key, err := conf.GetConfig("sys", "key")
		if err != nil {
			log.Fatal(err)
		}

		http.HandleFunc("/", httpHandler)
		http.ListenAndServeTLS(addr, cert, key, nil)
	case "tcp":
	default:
		log.Fatal("Please config your proto")
	}
}
