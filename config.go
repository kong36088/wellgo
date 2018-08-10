package wellgo

import (
	"runtime"
	"flag"
	"github.com/larspensjo/config"
	"fmt"
	"os"
	"errors"
	"path/filepath"
	"sync"
)

var (
	conf       *Config
	curPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	appPath    = curPath
	configFile = flag.String("config", curPath+"config/config.ini", "General configuration file")
)

type Config struct {
	values *sync.Map
}

//topic list
//TODO 热加载config or 每次重新加载config file

func NewConfig() *Config {
	return &Config{
		values: &sync.Map{},
	}
}

func GetConfigInstance() *Config {
	if conf == nil {
		InitConfig()
	}
	return conf
}

func InitConfig() error {
	conf = NewConfig()

	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	cfgSecs, err := config.ReadDefault(*configFile)
	if err != nil {
		return errors.New(fmt.Sprintf("Fail to find %s %s", *configFile, err))
	}

	for _, s := range cfgSecs.Sections() {
		options, err := cfgSecs.SectionOptions(s)
		if err != nil {
			logger.Error("Read options of file %s s %s  failed, %s\n", *configFile, s, err)
			continue
		}
		section := &sync.Map{}
		for _, v := range options {
			option, err := cfgSecs.String(s, v)
			if err != nil {
				logger.Error("Read file %s option %s failed, %s\n", *configFile, v, err)
				continue
			}
			section.Store(v, option)
		}
		conf.values.Store(s, section)
	}
	return nil
}

func (c *Config) GetConfig(section string, option string) (string, error) {
	var (
		sectionCfg interface{}
		optionCfg  interface{}
		found      bool
	)
	if sectionCfg, found = c.values.Load(section); !found {
		return "", errors.New(fmt.Sprintf("config section '%s' not found", section))
	}

	if optionCfg, found = sectionCfg.(map[string]string)[option]; !found {
		return "", errors.New(fmt.Sprintf("config option '%s' not found", option))
	}

	return optionCfg.(string), nil
}

func (c *Config) GetSection(section string) (sync.Map, error) {
	var (
		sectionCfg interface{}
		found      bool
	)
	if sectionCfg, found = c.values.Load(section); !found {
		return sync.Map{}, errors.New(fmt.Sprintf("config section '%s' not found", section))
	}

	return sectionCfg.(sync.Map), nil
}

func (c *Config) GetAllConfig() sync.Map {
	return *c.values
}
