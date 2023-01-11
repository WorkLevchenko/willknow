package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/WorkLevchenko/willknow/internal/app/apiserver"
)

var (
	configPath string
)

// Инициализирует конфиг, указывает путь до него
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// Главная функция, инициализирует конфиг и запускает apiserver.
func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
