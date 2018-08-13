package wellgo

import (
	"github.com/cihub/seelog"
)

var (
	logger seelog.LoggerInterface
)

func GetLoggerInstance() seelog.LoggerInterface {
	if logger == nil {
		InitLogger()
	}
	return logger
}

func NewLogger() (seelog.LoggerInterface, error) {
	return seelog.LoggerFromConfigAsFile(confPath + "seelog.xml")
}

func InitLogger() error {
	var err error

	if logger, err = seelog.LoggerFromConfigAsFile(confPath + "seelog.xml"); err != nil {
		return err
	}

	return nil
}

func CloseLogger() {
	logger.Flush()
	logger.Close()
}
