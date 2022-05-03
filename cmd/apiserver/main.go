package main

import (
	"flag"
	"fmt"
	"go-rest-api/internal/app/apiserver"
	"log"

	"github.com/BurntSushi/toml"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	fmt.Println(configPath)
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}