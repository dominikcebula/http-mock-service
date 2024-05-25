package main

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Address string
	}
	Rules []Rule
}

type Rule struct {
	Request  RequestRule
	Response ResponseRule
}

type RequestRule struct {
	Path string
}

type ResponseRule struct {
	Body string
}

func readConfig() (Config, error) {
	configFileContent, err := os.ReadFile("config.yaml")
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(configFileContent, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
