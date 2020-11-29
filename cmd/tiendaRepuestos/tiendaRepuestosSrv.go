package main

import (
	"Golang-seminar-UNICEN-2020/cmd/tiendaRepuestos/internal/config"
	"flag"
	"fmt"
	"os"
)

func main() {
	configFile := flag.String("config", "./config/config.yaml", "this is the service config")
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(cfg.DB.Driver)
	fmt.Println(cfg.Version)
}
