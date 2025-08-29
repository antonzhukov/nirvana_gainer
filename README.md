# Telegram Sunday Reminder Bot (Go)

A simple Telegram bot that sends a reminder message every Saturday at 22:00 Moscow time (2 hours before Sunday starts).

## Features

- 🚀 **Simple & Fast**: Built with Go 1.23 for optimal performance
- 🤖 **Telegram Integration**: Uses official Telegram Bot API
- ⏰ **Automated Scheduling**: Runs via GitHub Actions every Saturday
- 🐳 **Container Ready**: No file dependencies, pure environment variables
- 🔒 **Secure**: No sensitive data stored in files

## Requirements

- **Go**: 1.23 or later
- **Telegram Bot Token**: Get from [@BotFather](https://t.me/botfather)
- **Chat ID**: The Telegram chat/group where reminders will be sent

## Quick Start

### 1. Set Environment Variables

```bash
export TELEGRAM_BOT_TOKEN="your_bot_token_here"
export TELEGRAM_CHAT_ID="your_chat_id_here"
```

### 2. Build and Run

```bash
make build
make run
```

## Makefile Commands

- `make build`: Build the binary (`sunday-bot`)
- `make run`: Run in foreground (requires env vars to be set)
- `make start`: Start as a background daemon
- `make stop`: Stop the daemon
- `make status`: Show daemon status
- `make logs`: Tail the log file
- `make clean`: Remove build artifacts

## Configuration

### Environment Variables

The bot requires these environment variables:

| Variable | Description | Example |
|----------|-------------|---------|
| `TELEGRAM_BOT_TOKEN` | Your Telegram bot token | `123456789:ABCdefGHIjklMNOpqrsTUVwxyz` |
| `TELEGRAM_CHAT_ID` | Target chat ID | `123456789` |

### How to Get These Values

1. **Bot Token**: 
   - Message [@BotFather](https://t.me/botfather) on Telegram
   - Create a new bot with `/newbot`
   - Copy the token provided

2. **Chat ID**:
   - Add your bot to a chat/group
   - Send a message in the chat
   - Visit: `https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates`
   - Look for the `chat.id` field

## GitHub Actions Workflow

The bot automatically runs every Saturday at 22:00 Moscow time via GitHub Actions.

### Setup

1. **Add Repository Secrets**:
   - Go to your repo → Settings → Secrets and variables → Actions
   - Add `TELEGRAM_BOT_TOKEN` and `TELEGRAM_CHAT_ID`

2. **Workflow Features**:
   - **Schedule**: Every Saturday at 22:00 Moscow time (19:00 UTC)
   - **Manual Trigger**: Available via `workflow_dispatch`
   - **Go Version**: Uses latest Go 1.23
   - **Caching**: Optimized Go module caching

### Manual Execution

You can manually trigger the workflow:
1. Go to **Actions** tab in your repository
2. Click **Sunday Reminder Bot**
3. Click **Run workflow**

## Local Development

### Run with Go directly

```bash
export TELEGRAM_BOT_TOKEN="your_token"
export TELEGRAM_CHAT_ID="your_chat_id"
go run .
```

### Run with Makefile

```bash
export TELEGRAM_BOT_TOKEN="your_token"
export TELEGRAM_CHAT_ID="your_chat_id"
make run
```

## Docker Support

```bash
docker run -e TELEGRAM_BOT_TOKEN="token" \
           -e TELEGRAM_CHAT_ID="id" \
           your-image
```

## Project Structure

```
tgbot/
├── main.go                 # Main bot logic
├── Makefile               # Build and run commands
├── go.mod                 # Go module definition (Go 1.23)
├── .github/workflows/     # GitHub Actions automation
│   └── sunday-reminder.yml
└── README.md              # This file
```

## Dependencies

- **Go**: 1.23+
- **Telegram Bot API**: `github.com/go-telegram-bot-api/telegram-bot-api/v5`

## Notes

- The bot sends a single reminder message and then exits
- Designed to run via cron/GitHub Actions for scheduled execution
- No persistent storage or complex scheduling logic
- Environment variables must be set before running

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test with `make build && make run`
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE). 