package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

var (
	ServerCfg cfgs
)

type cfgs struct {
	file string

	WebdavPort    string
	DwnpicsrvPort string
}

func LoadCfg() error {
	return ServerCfg.load()
}

func (cf *cfgs) load() (err error) {
	cf.file = *flag.String("c", "../cfg/cfg.json", "the config file")

	// 读取配置文件
	f, err := os.Open(cf.file)
	if err != nil {
		return err
	}

	defer f.Close()

	c, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(c, &cf)
	if err != nil {
		return err
	}

	return nil

}