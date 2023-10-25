package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Servers []struct{ Address string }
}

// jak returnowac error jak sa dwa returny? pointer?
func readConfig() (*Config, error) {
	var config *Config

	b, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("err reading config: %s", err)
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return nil, fmt.Errorf("err unmarshalling: %s", err)
	}

	return config, nil
}
