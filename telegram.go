package main

import (
	"log"
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func createBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(EnvConfig.Token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func sendMessage(tbot *tgbotapi.BotAPI, chatID int64, msgTxt string) {
	msg := tgbotapi.NewMessage(chatID, msgTxt)

	tbot.Send(msg)
}

func checkUpdate(update *tgbotapi.Update) bool {
	isText := reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != ""

	return isText
}

func startBot() *tgbotapi.BotAPI {
	bot := createBot()

	sendMessage(bot, EnvConfig.AdminId, "I am online!")

	return bot
}

func TelegramBot() {
	tgBot := startBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgBot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if checkUpdate(&update) {
			var msgTest string

			log.Println(update.Message.From.ID)

			switch update.Message.Text {
			case "/start":
				msgTest = "Welcome to Bot"
				chatID := update.Message.Chat.ID
				log.Printf("Bot added to chat with ID: %d\n", chatID)

			default:
				msgTest = "This is test version. Nice to meet you"
			}

			if update.Message.From.ID == EnvConfig.AdminId {
				msgTest = msgTest + " " + "[ADMIN MODE]"
			}

			sendMessage(tgBot, update.Message.Chat.ID, msgTest)
		}
	}

}
