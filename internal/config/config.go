package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Env    string `yaml:"env"`
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type DB struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"ssl_mode"`
}

func Load() *Config {
	var cfg Config

	err := cleanenv.ReadConfig("config.yml", &cfg)

	if err != nil {
		log.Fatalf("error while read config: %v", err)
	}

	return &cfg
}
