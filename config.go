package main

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	Server struct {
		Host string
		Port int
	}
	Rules []Rule
}

type Rule struct {
	Request  RequestRule
	Response ResponseRule
}

type RequestRule struct {
	Path   string
	Method string
}

type ResponseRule struct {
	Code    int
	Delay   time.Duration
	Headers Headers
	Body    string
}

type Headers map[string]string

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
