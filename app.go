package wellgo

import "log"

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
}
