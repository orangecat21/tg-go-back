package bot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init() error {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN not set")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60 // таймаут long polling

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		text := update.Message.Text

		// Примитивный эхо-бот
		msg := tgbotapi.NewMessage(chatID, "Ты написал: "+text)

		// Пример реакции на команду /start
		if text == "/start" {
			msg.Text = "Привет! Я Go-бот 🤖. Напиши мне что-нибудь."
		}

		if _, err := bot.Send(msg); err != nil {
			log.Println("send error:", err)
		}
	}

	return nil
}
