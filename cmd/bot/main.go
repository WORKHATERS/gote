package main

import (
	"gote/internal/bot"
	c "gote/internal/commands"
	h "gote/internal/handlers"
	"gote/internal/utils/env"
	"os"
)

func main() {
	_ = env.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Токен отсутствует")
	}

	b := bot.NewBot(token)

	commands := c.NewCommands()
	commands.Add("/start", StartHandler)
	b.WithCommands(&commands)

	handlers := h.NewHandlers()
	handlers.Add(h.Message, MessageHandler)
	b.WithHandlers(&handlers)

	stateMachine := createStateMachine()

	b.WithState(stateMachine)

	b.RunUpdate()
}
