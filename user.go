package main

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TelegramID    int64 `gorm:"uniqueIndex"`
	Name          *string
	Birthday      *time.Time
	DataCompleted bool `gorm:"default: false"`
}

type UserParams struct {
	TelegramID    int64
	Name          *string
	Birthday      *time.Time
	DataCompleted *bool
}

func CheckUser(telegramID int64) *User {
	var dbUser User

	result := Db.Where(User{TelegramID: telegramID}).First(&dbUser)

	if result.RowsAffected == 0 {
		return nil
	}

	return &dbUser
}

func CreateUser(userParams *UserParams) *User {
	dbUser := User{
		TelegramID: userParams.TelegramID,
	}

	result := Db.Create(&dbUser)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return &dbUser
}

func GetUserBirthday(telegramID int64) (*User, error) {
	var dbUser User

	result := Db.Where(User{TelegramID: telegramID}).Not("birthday", nil).First(&dbUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &dbUser, nil
}

func UpdateUserBirthday(telegramID int64, birthday time.Time) (*User, error) {
	var dbUser User

	//result := Db.Where(User{TelegramID: telegramID}).First(&dbUser)

	updateData := map[string]interface{}{
		"Birthday":      birthday,
		"DataCompleted": true,
	}

	result := Db.Model(&dbUser).Where(User{TelegramID: telegramID}).Updates(updateData)

	if result.Error != nil {
		return nil, result.Error
	}

	return &dbUser, nil
}