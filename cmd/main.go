package main

import (
	"github.com/BurntSushi/toml"
	"github.com/ygjken/sender/pkg/webhook"
)

type Conf struct {
	Connect connect
}

type connect struct {
	Url string `toml:"url"`
}

func main() {
	var config *Conf

	_, err := toml.DecodeFile("./config.toml", &config)
	if err != nil {
		panic("Can't read config.toml")
	}

	webhook.WebHook(config.Connect.Url)
}
