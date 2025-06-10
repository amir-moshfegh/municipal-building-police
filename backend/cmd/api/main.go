package main

import (
	"fmt"
	"github.com/amir-moshfegh/municipal-building-police/config"
	"github.com/amir-moshfegh/municipal-building-police/internal/adapter/http"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to get config: %v", err))
	}

	server := http.NewServer(cfg)

	http.InitRoutes(server, cfg)
	http.InitSwaggerRoutes(server, cfg.Server)

	if err := server.Start(cfg.Server.Port); err != nil {
		panic(fmt.Sprintf("failed to create server: %v", err))
	}
}
