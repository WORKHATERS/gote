package main

import (
	"context"
	"fmt"
	gotebot "gote/pkg/bot"
	"gote/pkg/types"
)

func RequestName(ctx context.Context, update *types.Update, bot *gotebot.Bot) {
	myService, ok := gotebot.Resolve[MyService](bot.Dependencies)
	if !ok {
		fmt.Println("Не нашел сервис")
	}

	myService.SayHello()
	id := update.Message.Chat.Id
	fmt.Println("Введите имя:")
	bot.State.SetState(id, "writeNameState")
}

func WriteName(ctx context.Context, update *types.Update, bot *gotebot.Bot) {
	id := update.Message.Chat.Id
	text := update.Message.Text
	bot.Store.AddData(id, "name", text)
	nameFromStore := bot.Store.GetData(id, "name")
	fmt.Println("Из ханилища:", nameFromStore)
	bot.State.SetState(id, "requestMailState")
	RequestMail(ctx, update, bot)
}

func RequestMail(ctx context.Context, update *types.Update, bot *gotebot.Bot) {
	fmt.Println("Введите почту:")

	id := update.Message.Chat.Id
	bot.State.SetState(id, "writeMailState")
}

func WriteMail(ctx context.Context, update *types.Update, bot *gotebot.Bot) {
	fmt.Println("Почта записана")
}
