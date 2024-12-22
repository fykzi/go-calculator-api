package config

import (
	"flag"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
    Host string `yaml:"host" env-default:"127.0.0.1"`    
    Port int `yaml:"port" env-default:"8000"`
    LogLevel string `yaml:"log_level" env-default:"INFO"`
}

func LoadConfig() *Config {
    var pathToConfig string
    var config Config

    flag.StringVar(&pathToConfig, "config", "./config/config.yaml", "path to config file")
    flag.Parse()

    if err := cleanenv.ReadConfig(pathToConfig, &config); err != nil {
        panic("cannot read config: " + err.Error())
    }

    return &config
}
