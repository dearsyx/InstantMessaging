package config

import "gopkg.in/ini.v1"

type Email struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Addr     string `ini:"addr"`
}

type Mongo struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type SystemConfig struct {
	Email Email `ini:"email"`
	Mongo Mongo `ini:"mongo"`
}

var Config SystemConfig

func LoadConfig(path string) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return err
	}
	err = cfg.MapTo(&Config)
	if err != nil {
		return err
	}
	return nil
}
