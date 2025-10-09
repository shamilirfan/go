package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configaretions Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int64
}

func loadConfig() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed to load the env variables", err)
	}

	version := os.Getenv("VERSION")

	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")

	if serviceName == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)

	if err != nil {
		fmt.Println("Port must be number")
	}

	configaretions = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
	}
}

func GetConfig() Config {
	loadConfig()
	return configaretions
}
