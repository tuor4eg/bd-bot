package main

import (
	"fmt"
	"log"
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var tgBot *tgbotapi.BotAPI

func createBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(EnvConfig.Token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = EnvConfig.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func sendMessage(chatID int64, msgTxt string) {
	msg := msgTxt

	if chatID == EnvConfig.AdminId {
		msg = msg + " [ADMIN MODE]"
	}

	sendMsg := tgbotapi.NewMessage(chatID, msg)

	if _, err := tgBot.Send(sendMsg); err != nil {
		log.Panic(err)
	}
}

func checkUpdate(update *tgbotapi.Update) bool {
	isText := reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" && update.Message != nil

	return isText
}

func startBot() {
	tgBot = createBot()

	sendMessage(EnvConfig.AdminId, "I am online!")
}

func TelegramBot() {
	startBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgBot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			go callbacks(update)
		} else if update.Message.IsCommand() {
			go commands(update)
		} else {
			go messages(update)
		}
	}
}

func callbacks(update tgbotapi.Update) {

}

func commands(update tgbotapi.Update) {
	command := update.Message.Command()

	var msg string

	chatID := update.Message.Chat.ID

	switch command {
	case Commands.Start:
		msg = commandStart(chatID)
	case Commands.Birthday:
		msg = commandBirthday(chatID)
	default:
		msg = "This is test version. Nice to meet you"
	}

	sendMessage(chatID, msg)
}

func messages(update tgbotapi.Update) {
	var msg string
	var state State
	var err error

	chatID := update.Message.From.ID

	state, err = GetState(chatID)

	if err == nil {
		switch state.Command {
		case Commands.Birthday:
			msg = messageBirthday(chatID, update.Message.Text)
		}
	} else {
		msg = Errors.INTERNAL_ERROR
	}

	sendMessage(chatID, msg)
}

/**
* Commands
 */

func commandStart(chatID int64) string {
	var user *User

	msg := "Nice to see you again!"

	user, _ = CheckUser(chatID)

	if user == nil {
		user = CreateUser(&UserParams{TelegramID: chatID})

		msg = "Welcome to Birthday bot. Nice to meet you."
	}

	return msg
}

func commandBirthday(chatID int64) string {
	msg := "Enter your birthday date"

	user, _ := GetUserBirthday(chatID)

	if user != nil {
		msg = fmt.Sprintf("Your birthday date is %s. Enter new date if you want to change it.", user.Birthday.Format("02-01-2006"))
	}

	SetState(chatID, Commands.Birthday, Statuses.Pending)

	return msg
}

/**
* Messages
 */

func messageBirthday(chatID int64, date string) string {
	msg := SetUserBirthday(chatID, date)

	ClearState(chatID)

	return msg
}
