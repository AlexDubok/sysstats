package config

import (
	"flag"
	"fmt"
	"os"
)

// AppConfig structure to hold your configuration values
type AppConfig struct {
	Port string `json:"port"`
}

var Config AppConfig

// LoadConfig initializes the configuration from CLI arguments and environment variables
func LoadConfig() *AppConfig {
	port := "8080"

	flag.StringVar(&port, "port", getEnv("SYSSTATS_PORT", port), "HTTP server port")

	flag.Parse()
	// save port as Appconfig.Port
	Config.Port = port

	return &Config
}

// getEnv gets a value from the environment or uses a default
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetPort returns the port to run the server on
func GetPort() string {
	return fmt.Sprintf(":%s", Config.Port)
}
