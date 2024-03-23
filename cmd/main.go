package cmd

import (
	"article-service/internal/app"
	"article-service/internal/config"
	"flag"
)

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
