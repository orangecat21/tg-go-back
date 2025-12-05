package bot

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// UserData структура для хранения данных пользователя от Telegram Mini App
type UserData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName,omitempty"`
	Username  string `json:"username,omitempty"`
	PhotoURL  string `json:"photoUrl,omitempty"`
}

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

	// Запуск HTTP сервера для обработки данных от Telegram Mini App
	go func() {
		http.HandleFunc("/user-data", handleUserData)
		log.Println("HTTP server started on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("HTTP server error:", err)
		}
	}()

	return nil
}

// handleUserData обработчик для получения данных пользователя от Telegram Mini App
func handleUserData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var userData UserData
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		log.Printf("Error decoding user data: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Логируем полученные данные
	log.Printf("Received user data: %+v", userData)

	// Отправляем ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
