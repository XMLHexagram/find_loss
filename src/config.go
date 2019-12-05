package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type DBConfig struct {
	Address  string
	User     string
	Password string
	DBName   string
}

type webConfig struct {
	Port string
}

type Conf struct {
	DB  DBConfig
	Web webConfig
}

func (s *Service) ConfigInit() {
	confPath := "./FL_config/conf.toml"
	//c := new(Conf)

	_, err := toml.DecodeFile(confPath, &s.Config)
	DealError(err)
	fmt.Println(s.Config)
	//s.Config = *c
}
