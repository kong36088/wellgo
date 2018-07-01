package wellgo

import (
	"github.com/alecthomas/log4go"
)

var (
	logger *Logger
)

type Logger struct {
	log4go.Logger
}

func GetLogger() *Logger {
	if logger == nil {
		logger = &Logger{}
	}
	return logger
}

func (logger *Logger) Init() error {
	logPath, err := conf.GetConfig("sys", "log_conf")
	if err != nil {
		return err
	}
	logger.LoadConfiguration(appPath + logPath)
	return OK
}

func (logger *Logger) CloseLog() {
	logger.Close()
}
