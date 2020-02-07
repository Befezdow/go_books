package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/befezdow/go-books-rest-api/internal/gateway"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config_path", "configs/gateway-config.toml", "Path to config file")
}

func main() {
	flag.Parse()
	var config = gateway.NewConfig()
	var _, err = toml.DecodeFile(configPath, config)
	fmt.Println(config)
	if err != nil {
		log.Fatal(err)
	}
	var gatewayInstance = gateway.NewGateway(config)
	if err := gatewayInstance.Start(); err != nil {
		log.Fatal(err)
	}
}
