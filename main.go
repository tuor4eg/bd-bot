package main

import (
	"log"
)

var EnvConfig *Config

func main() {
	log.Println("Starting bot...")

	EnvConfig = LoadConfig()

	InitDb()
	TelegramBot()
}
