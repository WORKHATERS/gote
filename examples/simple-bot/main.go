package main

import (
	"context"
	"gote/internal/env"
	gotebot "gote/pkg/bot"
	"os"
)

const (
	startState       = "start"
	writeNameState   = "write_name"
	requestMailState = "request_mail"
	writeMailState   = "write_mail"
)

func main() {
	_ = env.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Токен отсутствует")
	}

	ctx, close := context.WithCancel(context.Background())
	defer close()

	bot := gotebot.NewBot(ctx, token)

	bot.OnState(startState, RequestName)
	bot.OnState(writeNameState, WriteName)
	bot.OnState(requestMailState, RequestMail)
	bot.OnState(writeMailState, WriteMail)

	bot.OnCommand("/start", startState)

	bot.Run()
}
