package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	err := gonfig.GetConf("echo-rest/config/config.json", &conf)

	fmt.Print(err)

	if err != nil {
		panic("No connection")
	}

	return conf
}