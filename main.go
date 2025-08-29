package main

import (
	"log"
	"os"
	"strconv"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}
	return v
}

func getMoscowTime() time.Time {
	// Moscow timezone
	moscowLoc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Printf("failed to load Moscow location, using fixed zone: %v", err)
		// Fallback to fixed zone if location loading fails
		moscowLoc = time.FixedZone("MSK", 3*60*60)
	}
	return time.Now().In(moscowLoc)
}

func getMessageByDay(moscowTime time.Time) string {
	weekday := moscowTime.Weekday()

	// Check if it's Sunday
	if weekday == time.Sunday {
		return "Началось воскресенье, а это значит, что хватит втыкать в телефон. Позволь себе немного цифрового детокса"
	}

	// Check if it's Monday
	if weekday == time.Monday {
		return "Начался понедельник, и только наш чатик поможет тебе не умереть со скуки в эту рабочую неделю"
	}

	return ""
}

func main() {
	// Get required environment variables
	botToken := mustEnv("TELEGRAM_BOT_TOKEN")
	chatIDStr := mustEnv("TELEGRAM_CHAT_ID")

	// Parse chat ID
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatalf("invalid TELEGRAM_CHAT_ID: %v", err)
	}

	// Create Telegram bot
	bot, err := telegram.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("failed creating bot: %v", err)
	}
	bot.Debug = false

	// Get current Moscow time and determine message
	moscowTime := getMoscowTime()
	message := getMessageByDay(moscowTime)
	if message == "" {
		os.Exit(0)
		return
	}

	// Send message
	msg := telegram.NewMessage(chatID, message)
	if _, err := bot.Send(msg); err != nil {
		log.Fatalf("send failed: %v", err)
	}

	log.Printf("message sent at %s: %s", moscowTime.Format("15:04"), message)
}
