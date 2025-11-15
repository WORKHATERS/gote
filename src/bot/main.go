package main

import (
	"context"
	"fmt"
	"gote/internal/bot"
	"gote/pkg/types"
)

func main() {
	token := ""
	if token == "" {
		panic("Токен отсутствует")
	}

	bot := bot.NewBot(token)
	ctx := context.Background()
	user, err := bot.GetMe(ctx, types.GetMe{})
	if err != nil {
		fmt.Println("Ошибка GetMe:", err)
	}

	fmt.Printf("Бот: %s (@%s), ID: %d\n", user.FirstName, user.Username, user.Id)
}
