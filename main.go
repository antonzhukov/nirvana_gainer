package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

// loadEnvFile loads environment variables from .env file in the root directory
func loadEnvFile() error {
	// Try to find .env file in current directory first
	if err := godotenv.Load(); err == nil {
		log.Println("Loaded .env file from current directory")
		return nil
	}

	// If not found, try to find it in the root directory (where main.go is located)
	execPath, err := os.Executable()
	if err != nil {
		log.Printf("Could not determine executable path: %v", err)
	} else {
		rootDir := filepath.Dir(execPath)
		envPath := filepath.Join(rootDir, ".env")
		if err := godotenv.Load(envPath); err == nil {
			log.Printf("Loaded .env file from: %s", envPath)
			return nil
		}
	}

	// Try to load from project root (assuming we're running from project directory)
	projectRoot, err := os.Getwd()
	if err == nil {
		envPath := filepath.Join(projectRoot, ".env")
		if err := godotenv.Load(envPath); err == nil {
			log.Printf("Loaded .env file from project root: %s", envPath)
			return nil
		}
	}

	log.Println("No .env file found, using system environment variables")
	return nil
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}
	return v
}

func main() {
	// Load environment variables from .env file
	if err := loadEnvFile(); err != nil {
		log.Printf("Warning: could not load .env file: %v", err)
	}

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
