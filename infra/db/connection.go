package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(config *config.DBConfig) string {
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s ",
		config.User,
		config.Passwrod,
		config.Host,
		config.Port,
		config.Name,
	)

	if !config.EnableSSLMODE {
		connectionString += "sslmode=disable"
	} else {
		connectionString += "sslmode=require"
	}

	return connectionString
}

func NewConnection(config *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(config)
	dbCon, err := sqlx.Connect("postgres", dbSource)

	if err != nil {
		fmt.Println("❌ Database connection error:", err)
		return nil, err
	}

	fmt.Println("✅ Database connected successfully!")
	return dbCon, nil
}
