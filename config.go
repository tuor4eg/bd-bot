package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds the configuration settings for the application.
type Config struct {
	Environment string
	Token       string
	AdminId     int64
	Debug       bool
	DbHost      string
	DbPort      int64
	DbUser      string
	DbPassword  string
	DbName      string
}

// LoadConfig loads the configuration from environment variables.
func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development" // Default to development if not set
	}

	token := os.Getenv("TOKEN")

	adminId, err := strconv.ParseInt(os.Getenv("ADMIN_ID"), 10, 64)

	if err != nil {
		log.Println("No admin ID found.", err)
	}

	if len(token) == 0 {
		log.Panic("Telegram bot token is empty.")
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))

	if err != nil {
		log.Println("No debug option. Set to TRUE")

		debug = true
	}

	dbHost := os.Getenv("DB_HOST")

	if len(dbHost) == 0 {
		log.Panic("Database host is empty.")
	}

	dbPort, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)

	if err != nil {
		log.Println("No database port found", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbName := os.Getenv("DB_NAME")

	if len(dbName) == 0 {
		log.Panic("Database name is empty.")
	}

	return &Config{
		Environment: env,
		Token:       token,
		AdminId:     adminId,
		Debug:       debug,
		DbHost:      dbHost,
		DbPort:      dbPort,
		DbUser:      dbUser,
		DbPassword:  dbPassword,
		DbName:      dbName,
	}
}
