package config

import (
	"os"
	"strings"
	"tungit/pkg/common/logger"

	"github.com/spf13/viper"
)

type ConfigAppEnv struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	Prefix      string `mapstructure:"PREFIX"`
	Host        string `mapstructure:"HOST"`
	Port        string `mapstructure:"PORT"`
}

func LoadConfig() *ConfigAppEnv {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // Enable reading environment variables

	loadEnvIntoViper()

	// If the .env file exists, read it
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logger.Error("Error reading config file: " + err.Error())
		}
	}

	config := &ConfigAppEnv{}
	err := viper.Unmarshal(config)
	if err != nil {
		logger.Error("Environment can't be loaded: " + err.Error())
	}

	// Set default config if empty
	setDefaultValues(config)

	return config
}

func setDefaultValues(config *ConfigAppEnv) {
	if config.Environment == "" {
		config.Environment = "development"
	}

	if config.Prefix == "" {
		config.Prefix = "api"
	}

	if config.Host == "" {
		config.Host = "localhost"
	}

	if config.Port == "" {
		config.Port = "8080"
	}
}

// Monkey patching for loading current env into viper
func loadEnvIntoViper() {
	// Retrieve the environment variables
	envVars := os.Environ()

	// Iterate over each environment variable
	for _, env := range envVars {
		// Split the string into key and value
		parts := strings.SplitN(env, "=", 2)
		key := parts[0]
		value := parts[1]

		// Print the key and value
		viper.Set(key, value)
	}
}
