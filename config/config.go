package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configaretions *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Passwrod      string
	EnableSSLMODE bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int64
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {
	err_1 := godotenv.Load()
	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort := os.Getenv("HTTP_PORT")
	port, err_2 := strconv.ParseInt(httpPort, 10, 64)
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	host := os.Getenv("DB_HOST")
	db_httpPort := os.Getenv("DB_PORT")
	db_port, err_3 := strconv.ParseInt(db_httpPort, 10, 64)
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	ssl_mode := os.Getenv("DB_ENABLE_SSL_MODE")
	enable_ssl_mode, err_4 := strconv.ParseBool(ssl_mode)

	if err_1 != nil {
		fmt.Println("Failed to load the env variables", err_1)
	}

	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	if serviceName == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}

	if httpPort == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}

	if err_2 != nil {
		fmt.Println("Port must be number")
	}

	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key is required")
		os.Exit(1)
	}

	if host == "" {
		fmt.Println("Host is required")
		os.Exit(1)
	}

	if err_3 != nil {
		fmt.Println("Database port must be number")
	}

	if db_httpPort == "" {
		fmt.Println("Http database port is required")
		os.Exit(1)
	}

	if name == "" {
		fmt.Println("Name is required")
		os.Exit(1)
	}

	if user == "" {
		fmt.Println("User is required")
		os.Exit(1)
	}

	if password == "" {
		fmt.Println("User is required")
		os.Exit(1)
	}

	if err_4 != nil {
		fmt.Println("Invalid enable ssl mode")
		os.Exit(1)
	}

	db_config := &DBConfig{
		Host:          host,
		Port:          int(db_port),
		Name:          name,
		User:          user,
		Passwrod:      password,
		EnableSSLMODE: enable_ssl_mode,
	}

	configaretions = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     port,
		JwtSecretKey: jwtSecretKey,
		DB:           db_config,
	}
}

func GetConfig() *Config {
	if configaretions == nil {
		loadConfig()
	}

	return configaretions
}
