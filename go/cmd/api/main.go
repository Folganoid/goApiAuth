package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"goApiAuth/go/internal/api"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/api.toml", "path to config path")
}

func main() {
	flag.Parse()

	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := api.Start(config); err != nil {
		log.Fatal(err)
	}
}