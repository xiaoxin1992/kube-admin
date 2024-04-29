package config

import (
	"github.com/BurntSushi/toml"
	"strings"
)

var conf = &Config{}

type Config struct {
	Level   string     `toml:"level"`
	Http    http       `toml:"http"`
	Logger  logger     `toml:"logger"`
	Cluster []*Cluster `toml:"cluster"`
	Jwt     jwt        `toml:"jwt"`
	MySQL   MySQL      `toml:"mysql"`
}

func Init(filename string) (err error) {
	_, err = toml.DecodeFile(filename, &conf)
	return
}

func GetConfig() *Config {
	return conf
}

func GetCluster(name string) *Cluster {
	for _, c := range conf.Cluster {
		if strings.TrimSpace(c.Name) == name {
			return c
		}
	}
	return nil
}
