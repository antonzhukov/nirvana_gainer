# Telegram Sunday Reminder Bot (Go)

Sends a Telegram message 2 hours before Sunday (local time), every week.

## Configuration

Use a `.env` file (recommended). See `.env.example`:

```env
TELEGRAM_BOT_TOKEN=...
TELEGRAM_CHAT_ID=...
PORT=8080
```

## Makefile usage

- `make build`: Build the binary (`sunday-bot`)
- `make run`: Run in foreground, loading vars from `.env`
- `make start`: Start as a background daemon, writing `sunday-bot.pid` and logs to `sunday-bot.log`
- `make stop`: Stop the daemon via PID file
- `make status`: Show whether the daemon is running
- `make logs`: Tail the log file

Variables:

- `ENV_FILE`: path to env file (default `.env`)
- `PID_FILE`: pid file path (default `sunday-bot.pid`)
- `LOG_FILE`: log file path (default `sunday-bot.log`)

## Run locally without Makefile

```bash
export TELEGRAM_BOT_TOKEN="<your-token>"
export TELEGRAM_CHAT_ID="<chat-id>"

go run .
```

## Notes

- The scheduler computes the next occurrence of Sunday 00:00 in your machine's local timezone and triggers 2 hours earlier.
- Health endpoint at `/:8080/healthz` by default. 