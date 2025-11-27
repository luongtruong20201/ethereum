package util

import (
	"flag"
	"fmt"
	"os"

	"github.com/rakyll/globalconf"
)

type ConfigManager struct {
	Db       Database
	ExecPath string
	Debug    bool
	Diff     bool
	DiffType string
	Paranoia bool
	conf     *globalconf.GlobalConf
}

var Config *ConfigManager

func ReadConfig(ConfigFile string, Datadir string, EnvPrefix string) *ConfigManager {
	if Config == nil {
		_, err := os.Stat(ConfigFile)
		if err != nil && os.IsNotExist(err) {
			fmt.Printf("config file '%s' doesn't exist, creating it\n", ConfigFile)
			os.Create(ConfigFile)
		}
		g, err := globalconf.NewWithOptions(&globalconf.Options{
			Filename:  ConfigFile,
			EnvPrefix: EnvPrefix,
		})
		if err != nil {
			fmt.Println(err)
		} else {
			g.ParseAll()
		}
		Config = &ConfigManager{ExecPath: Datadir, Debug: true, conf: g, Paranoia: true}
	}
	return Config
}

func (c *ConfigManager) Save(key string, value interface{}) {
	f := &flag.Flag{Name: key, Value: newConfValue(value)}
	c.conf.Set("", f)
}

func (c *ConfigManager) Delete(key string) {
	c.conf.Delete("", key)
}

type confValue struct {
	value string
}

func newConfValue(value interface{}) *confValue {
	return &confValue{fmt.Sprintf("%v", value)}
}

func (c confValue) String() string {
	return c.value
}

func (c confValue) Set(s string) error {
	c.value = s
	return nil
}
