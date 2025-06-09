package main

import (
	"fmt"
	"github.com/amir-moshfegh/municipal-building-police/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.Server.Host)
}
