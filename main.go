package main

import (
	"log"
	"os"
	"strconv"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}
	return v
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

	// Send reminder message
	msg := telegram.NewMessage(chatID, "Reminder: Sunday starts in 2 hours ⏰")
	if _, err := bot.Send(msg); err != nil {
		log.Fatalf("send failed: %v", err)
	}

	log.Println("reminder sent; exiting")
}
