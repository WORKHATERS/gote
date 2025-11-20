package main

import (
	"context"
	"fmt"
	gotebot "gote/internal/bot"
	"gote/internal/utils/env"
	"gote/pkg/methods"
	"gote/pkg/types"
	"os"
)

func main() {
	_ = env.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Токен отсутствует")
	}

	// ctx := ctx.CustomContext{
	// 	Token:     token,
	// 	GoContext: context.Background(),
	// }
	ctxClose, closeFunc := context.WithCancel(context.Background())
	ctx := context.WithValue(ctxClose, "token", token)

	bot := gotebot.NewBot(ctx)

	me, _ := methods.GetMe(ctx, types.GetMe{})
	fmt.Println(me.FirstName, me.Id, me.Username)

	bot.OnCommand("/start", StartHandler)

	stateMachine := createStateMachine()
	bot.WithState(stateMachine)

	bot.RunUpdate()
	closeFunc()
}
