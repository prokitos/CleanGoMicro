package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type MainConfig struct {
	Server     ServerConfig   `yaml:"server"`
	PostgresDB PostgresConfig `yaml:"postgres"`
	MongoDB    MongoConfig    `yaml:"mongo"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}
type PostgresConfig struct {
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Host   string `yaml:"host"`
	Name   string `yaml:"name"`
	Reload bool   `yaml:"reload"`
	Port   string `yaml:"port"`
}
type MongoConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (cfg *MainConfig) ConfigMustLoad(name string) {

	path := "./config/" + name + ".yaml"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("nothing from this path")
	}

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		panic("failed to read")
	}

}
