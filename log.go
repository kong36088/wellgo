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

func GetLoggerInstance() *Logger {
	if logger == nil {
		logger = &Logger{}
	}
	return logger
}

func (logger *Logger) Init() error {
	conf := NewConfig()
	logPath, err := conf.Get("config", "log", "path")
	if err != nil {
		return err
	}
	logger.LoadConfiguration(appPath + logPath)
	return nil
}

func (logger *Logger) Close() {
	logger.Logger.Close()
}
