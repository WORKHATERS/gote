package main

import (
	"context"
	"fmt"
	"gote/pkg/bot"
	"gote/pkg/types"
)

func RequestName(ctx context.Context, update *types.Update, bot *bot.Bot) {
	id := update.Message.Chat.Id
	fmt.Println("Введите имя:")
	bot.SetState(id, writeNameState)
}

func WriteName(ctx context.Context, update *types.Update, bot *bot.Bot) {
	id := update.Message.Chat.Id
	fmt.Println("Имя записано")

	bot.SetState(id, requestMailState)
	RequestMail(ctx, update, bot)
}

func RequestMail(ctx context.Context, update *types.Update, bot *bot.Bot) {
	fmt.Println("Введите почту:")

	id := update.Message.Chat.Id
	bot.SetState(id, writeMailState)
}

func WriteMail(ctx context.Context, update *types.Update, bot *bot.Bot) {
	fmt.Println("Почта записана")
}
