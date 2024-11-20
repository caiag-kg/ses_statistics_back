package config

// Package config is responsible for handling application configuration.
// It provides functionality to read and parse configuration from a YAML file.

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"ses_back/internal/models"
)

// Config holds the application configuration.
var Config models.AppConfig

// GetConfigPath returns the path to the configuration file.
func GetConfigPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	configPath := filepath.Join(wd, "config.yaml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Panic(err)
	}

	return configPath, nil
}

// ReadConfig reads and parses the configuration from the YAML file.
func ReadConfig() {
	YamlPath, err := GetConfigPath()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(YamlPath)

	if err != nil {
		log.Panic(err)
	}

	err = yaml.Unmarshal(file, &Config)

	if err != nil {
		log.Panic(err)
	}
}