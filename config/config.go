package config

import "time"

type (
	MainConfig struct {
		Server   ServerConfig   `yaml:"server"`
		Database DatabaseConfig `yaml:"database"`
	}

	ServerConfig struct {
		Host    string        `yaml:"host"`
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}

	DatabaseConfig struct {
		Dsn string `yaml:"dsn"`
	}
)
