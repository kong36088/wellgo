package wellgo

import (
	"github.com/larspensjo/config"
	"fmt"
	"os"
	"errors"
	"path/filepath"
	"sync"
	"github.com/emirpasic/gods/sets/treeset"
)

var (
	curPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	appPath    = curPath
)

// threadsafe config
type Config struct {
	files  *treeset.Set
	values *sync.Map
}

//topic list
//TODO 热加载config or 每次重新加载config file

func NewConfig() *Config {
	return &Config{
		files:  treeset.NewWithStringComparator(),
		values: &sync.Map{},
	}
}

func (c *Config) LoadConfig(configFile string) error {
	configFile = curPath + "config/" + configFile + ".ini"
	if c.files.Contains(configFile) {
		return nil
	}

	cfgSecs, err := config.ReadDefault(configFile)
	if err != nil {
		return errors.New(fmt.Sprintf("Fail to find %s %s", configFile, err))
	}

	fileConfigs := &sync.Map{}
	// parse config contents
	for _, s := range cfgSecs.Sections() {
		options, err := cfgSecs.SectionOptions(s)
		if err != nil {
			logger.Error("Read options of file %s s %s  failed, %s\n", configFile, s, err)
			continue
		}
		section := &sync.Map{}
		for _, v := range options {
			option, err := cfgSecs.String(s, v)
			if err != nil {
				logger.Error("Read file %s option %s failed, %s\n", configFile, v, err)
				continue
			}
			section.Store(v, option)
		}
		c.values.Store(s, section)
	}
	// record config file
	c.files.Add(configFile)
	c.values.Store(configFile, fileConfigs)

	return nil
}

func (c *Config) Get(file string, section string, option string) (string, error) {
	var (
		fileCfg    interface{}
		sectionCfg interface{}
		optionCfg  interface{}
		found      bool
	)
	if err := c.LoadConfig(file); err != nil {
		return "", err
	}

	if fileCfg, found = c.values.Load(file); !found {
		return "", errors.New(fmt.Sprintf("config file '%s' not found", file))
	}
	if sectionCfg, found = fileCfg.(*sync.Map).Load(section); !found {
		return "", errors.New(fmt.Sprintf("config section '%s' not found", section))
	}

	if optionCfg, found = sectionCfg.(*sync.Map).Load(option); !found {
		return "", errors.New(fmt.Sprintf("config option '%s' not found", option))
	}

	return optionCfg.(string), nil
}

func (c *Config) GetSection(file string, section string) (sync.Map, error) {
	var (
		fileCfg    interface{}
		sectionCfg interface{}
		found      bool
	)
	if err := c.LoadConfig(file); err != nil {
		return sync.Map{}, err
	}
	if fileCfg, found = c.values.Load(file); !found {
		return sync.Map{}, errors.New(fmt.Sprintf("config file '%s' not found", file))
	}
	if sectionCfg, found = fileCfg.(*sync.Map).Load(section); !found {
		return sync.Map{}, errors.New(fmt.Sprintf("config section '%s' not found", section))
	}

	return *(sectionCfg.(*sync.Map)), nil
}

func (c *Config) GetLoadedFiles() []string {
	ret := make([]string, c.files.Size())
	for i, f := range c.files.Values() {
		ret[i], _ = f.(string)
	}
	return ret
}

func (c *Config) GetAllConfig() sync.Map {
	return *c.values
}
