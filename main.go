package main

import (
	t "bd-bot/dist"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting bot...")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	t.TelegramBot()
}
