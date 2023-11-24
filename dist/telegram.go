package telegram

import (
	"fmt"
	"log"
	"os"
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TelegramBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			var msgTest string

			fmt.Println(update)

			switch update.Message.Text {
			case "/start":
				msgTest = "Welcome to Bot"
			default:
				msgTest = "This is test version. Nice to meet you"
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgTest)
			bot.Send(msg)
		}
	}

}
