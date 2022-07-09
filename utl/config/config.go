package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
}

type Server struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	Timeout int    `yaml:"timeout"`
}

type DB struct {
	Url string `yaml:"url"`
}

func LoadConfig(path string) *Config {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	config := new(Config)

	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	return config
}
