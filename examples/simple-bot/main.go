package main

import (
	"context"
	"fmt"
	"gote/internal/env"
	gb "gote/pkg/bot"
	"log"
	"os"
)

type MyService struct {
	Name string
}

func (ms MyService) SayHello() {
	fmt.Println("Hello,", ms.Name)
}

func main() {
	// получение ключа бота из файла .env
	_ = env.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Токен отсутствует")
	}

	// создание контекста
	ctx, close := context.WithCancel(context.Background())
	defer close()

	// создание бота
	bot := gb.NewBot(ctx, gb.Config{
		Token:        token,
		Limit:        100,
		Timeout:      50,
		WorkersCount: 100,
	})

	// добавление зависимостей (DI)
	deps := gb.NewDependencies()

	myService := MyService{"My Services"}
	deps.Provide(myService)

	bot.AddDependencies(deps)

	// создание состояний
	startState := bot.State.NewState("start", RequestName)
	_ = bot.State.NewState("writeNameState", WriteName)
	_ = bot.State.NewState("requestMailState", RequestMail)
	_ = bot.State.NewState("writeMailState", WriteMail)

	// рабора с командами
	bot.State.OnCommand("/start", startState)

	// запуск цикла обновлений
	err := bot.Run()
	if err != nil {
		log.Println(err)
	}
}
