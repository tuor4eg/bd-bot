package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitDb() {
	connectDB()
	migrateDB()
}

func connectDB() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		EnvConfig.DbHost, EnvConfig.DbPort, EnvConfig.DbUser, EnvConfig.DbPassword, EnvConfig.DbName)

	var err error

	level := logger.Silent

	if EnvConfig.Debug {
		level = logger.Info
	}
	Db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})

	if err != nil {
		log.Fatal("Connection to database failed:", err)
	}

	log.Println("Successfully connected to database " + EnvConfig.DbName)
}

func migrateDB() {
	Db.AutoMigrate(&User{})
}

func CloseConnection() {
	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}
