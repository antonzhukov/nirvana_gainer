# Simple Makefile to run the bot and manage it as a daemon

BIN := sunday-bot
PID_FILE ?= sunday-bot.pid
LOG_FILE ?= sunday-bot.log

.PHONY: all build run start stop status logs clean

all: build

build:
	go build -o $(BIN) .

run: build
	@echo "[run] using environment variables from shell"
	@./$(BIN)

start: build
	@echo "[start] using environment variables from shell"
	@if [ -f "$(PID_FILE)" ] && kill -0 $$(cat "$(PID_FILE)") 2>/dev/null; then \
		echo "already running with PID $$(cat \"$(PID_FILE)\")"; \
		exit 0; \
	fi
	@nohup ./$(BIN) >> $(LOG_FILE) 2>&1 & echo $$! > $(PID_FILE)
	@echo "started, PID $$(cat \"$(PID_FILE)\")"

stop:
	@echo "[stop]"
	@if [ -f "$(PID_FILE)" ]; then \
		PID=$$(cat "$(PID_FILE)"); \
		if kill -0 $$PID 2>/dev/null; then \
			kill $$PID; \
			echo "sent SIGTERM to $$PID"; \
		fi; \
		rm -f "$(PID_FILE)"; \
	else \
		echo "no PID file"; \
	fi

status:
	@if [ -f "$(PID_FILE)" ] && kill -0 $$(cat "$(PID_FILE)") 2>/dev/null; then \
		echo "running (PID $$(cat \"$(PID_FILE)\"))"; \
	else \
		echo "not running"; \
	fi

logs:
	@echo "[logs] tailing $(LOG_FILE)"
	@tail -f $(LOG_FILE)

clean:
	rm -f $(BIN) $(PID_FILE) 
