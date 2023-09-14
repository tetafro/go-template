package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents application configuration.
type Config struct {
	ListenAddr string `yaml:"listen_addr"`
}

// ReadConfig returns configuration populated from the config file.
func ReadConfig(file string) (Config, error) {
	data, err := os.ReadFile(file) //nolint:gosec
	if err != nil {
		return Config{}, fmt.Errorf("read file: %w", err)
	}
	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return Config{}, fmt.Errorf("unmarshal file: %w", err)
	}

	// Set defaults
	if conf.ListenAddr == "" {
		conf.ListenAddr = "0.0.0.0:8080"
	}

	return conf, nil
}
