package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DB   `json:"db"`
	Log  `json:"log"`
	Web  `json:"web"`
	GRPC `json:"grpc"`
}

type DB struct {
	Enable bool   `json:"enable"`
	DSN    string `json:"dsn"`
	Merge  bool   `json:"merge"`
}

type Log struct {
	Dir      string `json:"dir"`
	Level    string `json:"level"`
	MaxAge   int    `json:"max_age"`
	Duration int    `json:"duration"`
	MaxSize  int    `json:"max_size"`
}

type Web struct {
	Enable bool   `json:"enable"`
	Addr   string `json:"addr"`
}

type GRPC struct {
	Enable bool   `json:"enable"`
	Addr   string `json:"addr"`
}

var conf *Config

func LoadConfig(confPath string) error {
	confFile, err := os.Open(confPath)
	if err != nil {
		return fmt.Errorf("config path %s,err:%v", confPath, err)
	}
	defer confFile.Close()

	conf = &Config{}
	if err := json.NewDecoder(confFile).Decode(conf); err != nil {
		return err
	}
	return nil
}

func Conf() *Config {
	return conf
}
