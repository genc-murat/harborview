package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config struct defines the structure of the YAML configuration.
type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Registry struct {
		BaseURL  string `yaml:"baseUrl"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"registry"`
}

// Load reads the configuration from config/config.yaml.
func Load() Config {
	// Determine the executable's directory
	execDir, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable directory: %v", err)
	}
	configDir := filepath.Dir(execDir)

	// Construct the config file path
	configPath := filepath.Join(configDir, "", "config.yaml")

	// Open the config file
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}

	return cfg
}
