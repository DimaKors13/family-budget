package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env          string       `yaml:"env" env-default:"local"`
	DBAccessInfo DBAccessInfo `yaml:"db_access_info"`
	HTTPServer   HTTPServer   `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8090"`
	TimeOut     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
}

type DBAccessInfo struct {
	DBAddress  string `yaml:"storage_address" env-required:"true"`
	DBUser     string `yaml:"db_user" env-default:"postgres"`
	DBPassword string `yaml:"db_password" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("Config path is not defined.")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file on path \"%s\" is not exist.", configPath)
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("Cannot read config file on path \"%s\".", configPath)
	}

	return &config
}
