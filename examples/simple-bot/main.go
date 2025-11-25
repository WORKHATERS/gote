package main

import (
	"context"
	"fmt"
	"gote/internal/env"
	"gote/pkg/bot"
	"gote/pkg/core"
	"gote/pkg/router"
	"gote/pkg/types"
	"gote/pkg/updater"
	"os"
)

type MyService struct {
	Name string
}

func (ms MyService) SayHello() {
	fmt.Println("Hello,", ms.Name)
}

// func StartHandler(ctx context.Context, bot core.BotContext, update types.Update) error {
// 	fmt.Println("Старт!")
// 	return nil
// }

// func StartMiddleware(next router.HandlerFunc) router.HandlerFunc {
// 	return func(ctx context.Context, bot core.BotContext, update types.Update) error {
// 		bot.Logger().Println("Вызов обработчика")
// 		return next(ctx, bot, update)
// 	}
// }

func main() {
	// получение ключа бота из файла .env
	// BOT_TOKEN=токен_из_BotFather
	_ = env.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Токен отсутствует")
	}

	// создание контекста
	ctx, closeFunc := context.WithCancel(context.Background())
	defer closeFunc()

	// создание роутера
	r := router.New()

	// Добавляем обработчик команды /start
	r.Command("/start", func(ctx context.Context, b core.BotContext, update types.Update) error {
		msg := update.Message
		if msg != nil {
			b.Logger().Printf("Получена команда /start от пользователя %d", msg.Chat.Id)
		}
		return nil
	})

	// Добавляем обработчик на любое сообщение, содержащее "hello"
	r.TextContains("hello", func(ctx context.Context, b core.BotContext, update types.Update) error {
		msg := update.Message
		if msg != nil {
			b.Logger().Printf("Пользователь %d написал hello!", msg.Chat.Id)
		}
		return nil
	})

	// Добавляем middleware для логирования всех апдейтов
	r.Use(func(next router.HandlerFunc) router.HandlerFunc {
		return func(ctx context.Context, b core.BotContext, update types.Update) error {
			b.Logger().Printf("Новый апдейт: %+v", update)
			return next(ctx, b, update)
		}
	})
	// создание бота
	b := bot.New(ctx, token,
		bot.WithRouter(r),
	)

	// полдучение обновлений от Telegram
	poller := updater.NewPoller(b)
	updates := poller.Start()
	for u := range updates {
		// встроенная обработка через роутер
		b.Handle(u)

		// ручная обработка
		// if msg := u.Message; msg != nil {
		// 	fmt.Println(u.Message.Text)
		// }
	}
}
