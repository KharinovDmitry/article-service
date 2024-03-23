package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	ConnStr string `yaml:"conn_str"`
	Address string `yaml:"address"`
}

func MustLoadConfig(configPath string) Config {
	config := Config{}
	file, err := os.ReadFile(configPath)

	if err != nil {
		panic("Config not found!")
	}

	err = yaml.Unmarshal(file, config)

	if err != nil {
		panic("Could not unmarshal config correct.")
	}

	return config
}
