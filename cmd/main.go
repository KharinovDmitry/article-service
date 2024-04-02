package main

import (
	_ "article-service/docs"
	"article-service/internal/app"
	"article-service/internal/config"
	"flag"
)

// @title Articles Service
// @host localhost:8080
// @BasePath /

func main() {
	var configPath string
	flag.StringVar(&configPath, "path", "", "config path")
	flag.Parse()
	if configPath == "" {
		panic("Config path not found")
	}

	cfg := config.MustLoadConfig(configPath)
	app.MustRun(cfg)
}
