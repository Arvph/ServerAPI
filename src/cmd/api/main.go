package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/arvph/ServerAPI/internal/app/api"
)

var (
	configPath string = "configs/api.toml"
)

func init() {
	// getting the path to config file as argument
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file .toml format")
}

func main() {
	// initialization of variable config Path with received value
	flag.Parse()
	log.Println("It works!")
	// server instance initialization
	config := api.NewConfig()
	// Get new configs from .toml if the file exists
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Cannot find configs file. using default values:", err)
	}
	server := api.New(config)

	// api server start
	log.Fatal(server.Start())
}
