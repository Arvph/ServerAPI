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
	// адрес до конфигурационного файла
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file .toml format")
}

func main() {
	flag.Parse()
	log.Println("It works!")

	// инициализация настроек сервера
	config := api.NewConfig()
	// читаем из .toml, в котором могут быть новые конфигурационные данные
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Cannot find configs file. Using default values:", err)
	}
	server := api.New(config)

	log.Fatal(server.Start())
}
