package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	ConnStr          string `yaml:"conn_str"`
	Address          string `yaml:"address"`
	DriverName       string `yaml:"driver_name"`
	MigrationsDir    string `yaml:"migrations_dir"`
	TimeoutDbContext int    `yaml:"timeout_db_context"`
}

func MustLoadConfig(configPath string) Config {
	config := Config{}
	file, err := os.ReadFile(configPath)

	if err != nil {
		panic("Config not found!")
	}

	err = yaml.Unmarshal(file, &config)

	if err != nil {
		panic("Could not unmarshal config correct.")
	}

	return config
}
