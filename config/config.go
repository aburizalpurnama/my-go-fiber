package config

import "time"

type (
	MainConfig struct {
		Server ServerConfig `yaml:"server"`
	}

	ServerConfig struct {
		Host    string        `yaml:"host"`
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}
)
