package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	dp "github.com/araddon/dateparse"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var UserState = make(map[int64]State)

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
			callbacks(update)
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
		msg = "Nice to see you again!"

		user := CheckUser(chatID)

		if user == nil {
			user = CreateUser(&UserParams{TelegramID: chatID})

			msg = "Welcome to Birthday bot. Nice to meet you."
		}

		log.Printf("Bot added to chat with ID: %d\n", chatID)
	case Commands.Birthday:
		msg = "Enter your birthday date"

		user, _ := GetUserBirthday(chatID)

		if user != nil {
			msg = fmt.Sprintf("Your birthday date is %s. Enter new date if you want to change it.", user.Birthday.Format("02-01-2006"))
		}

		setUserState(chatID, Commands.Birthday, Statuses.Pending)
	default:
		msg = "This is test version. Nice to meet you"
	}

	sendMessage(chatID, msg)
}

func messages(update tgbotapi.Update) {
	var msg string

	chatID := update.Message.From.ID

	state, err := getUserState(chatID)

	if err == nil {
		switch state.Command {
		case Commands.Birthday:
			msg = setUserBirthday(chatID, update.Message.Text)

			clearUserState(chatID)
		}
	}

	sendMessage(chatID, msg)
}

func clearUserState(chatID int64) {
	delete(UserState, chatID)
}

func getUserState(chatID int64) (State, error) {
	status, stateOk := UserState[chatID]

	if !stateOk {
		return State{}, errors.New(Errors.NO_USER_STATE)
	}

	return status, nil
}

func setUserState(chatID int64, command string, status int) {
	_, userOk := UserState[chatID]

	if !userOk {
		UserState[chatID] = State{
			Command: command,
			Status:  status,
		}
	}
}

func setUserBirthday(chatID int64, date string) string {
	var msg string

	birthday, dateErr := dp.ParseAny(date)
	_, updErr := UpdateUserBirthday(chatID, birthday)

	if dateErr == nil && updErr == nil {
		msg = fmt.Sprintf("Your birthday date is %s", birthday.Format("02-01-2006"))
	} else {
		msg = Errors.INTERNAL_ERROR
	}

	return msg
}
