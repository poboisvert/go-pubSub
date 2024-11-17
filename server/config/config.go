package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Host          string
	Port          string
	User          string
	Password      string
	Database      string
	RedisHost     string
	RedisPassword string
}

var conf *Config

func GetConfig() *Config {
	if conf != nil {
		return conf
	}

	conf = &Config{}

	configContent, err := os.ReadFile(".env")
	if err != nil {
		fmt.Println("Error reading .env file: ", err)
		os.Exit(1)
	}
	envVariables := strings.Split(string(configContent), "\n")
	envMap := make(map[string]string)
	for _, variable := range envVariables {
		if strings.Contains(variable, "=") {
			splitVariable := strings.SplitN(variable, "=", 2) // Split only on the first '='
			key := strings.TrimSpace(splitVariable[0])
			value := strings.TrimSpace(strings.Trim(splitVariable[1], "\"")) // Trim quotes from value
			envMap[key] = value
		}
	}

	// Apply environment variables to config
	for k, v := range envMap {
		switch k {
		case "POSTGRES_HOST":
			conf.Host = v
		case "POSTGRES_PORT":
			conf.Port = v
		case "POSTGRES_USER":
			conf.User = v
		case "POSTGRES_PASSWORD":
			conf.Password = v
		case "POSTGRES_DB":
			conf.Database = v
		case "REDIS_HOST":
			conf.RedisHost = v
		case "REDIS_PASSWORD":
			conf.RedisPassword = v
		}
	}

	return conf
}

func init() {
	conf = GetConfig()
}
