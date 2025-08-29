package main

import (
	"log"
	"os"
	"strconv"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	SUNDAY_STARTS_TIME = "19:40"
	SUNDAY_OVER_TIME   = "19:42"
)

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}
	return v
}

func getMoscowTime() time.Time {
	// Moscow timezone (UTC+3)
	moscowLoc := time.FixedZone("MSK", 3*60*60)
	return time.Now().In(moscowLoc)
}

func getMessageByTime(moscowTime time.Time) string {
	currentTime := moscowTime.Format("15:04")

	// Check if it's 19:37
	// if currentTime == SUNDAY_STARTS_TIME {
	return "Sunday starts in 2 hours"
	// }

	// Check if it's 19:39
	if currentTime == SUNDAY_OVER_TIME {
		return "Sunday is over, chat again"
	}

	// Default message for other times
	return "Reminder: Sunday starts in 2 hours ⏰"
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
	message := getMessageByTime(moscowTime)

	// Send message
	msg := telegram.NewMessage(chatID, message)
	if _, err := bot.Send(msg); err != nil {
		log.Fatalf("send failed: %v", err)
	}

	log.Printf("message sent at %s: %s", moscowTime.Format("15:04"), message)
}
