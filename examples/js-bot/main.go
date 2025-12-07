package main

import (
	"context"
	"os"

	"github.com/WORKHATERS/gote/internal/env"
	"github.com/WORKHATERS/gote/pkg/core"
	"github.com/WORKHATERS/gote/pkg/types"
	"github.com/WORKHATERS/gote/pkg/updater"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

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

	// создание бота
	b := core.NewBot(ctx, token,
		core.WithHTTPClient(nil),
		core.WithLogger(nil),
	)

	registry := new(require.Registry)

	vm := goja.New()
	setScript(vm)

	registry.Enable(vm)
	console.Enable(vm)

	var jsHandler func(*types.Message)
	err := vm.ExportTo(vm.Get("handler"), &jsHandler)

	if err != nil {
		panic(err)
	}

	vm.Set("SendMessage", b.SendMessage)
	vm.Set("ctx", ctx)

	// полдучение обновлений от Telegram
	poller := updater.NewPoller(b)
	updates := poller.Start()
	for u := range updates {
		if cb := u.CallbackQuery; cb != nil {
			message, err := types.CastTo[types.Message](cb.Message)
			if err != nil {
				b.Logger().Error(err.Error())
			}

			b.SendMessage(ctx, types.SendMessage{
				ChatId: message.Chat.Id,
				Text:   "Вы выбрали: " + u.CallbackQuery.Data,
			})

		}

		if msg := u.Message; msg != nil {
			if msg.Text == "update_js" {
				setScript(vm)
				err := vm.ExportTo(vm.Get("handler"), &jsHandler)

				if err != nil {
					panic(err)
				}
			}
			jsHandler(u.Message)

			buttons := []types.InlineKeyboardButton{
				{
					Text:         "1",
					CallbackData: "1",
				},
				{
					Text:         "2",
					CallbackData: "2",
				},
			}

			b.SendMessage(ctx, types.SendMessage{
				ChatId: u.Message.Chat.Id,
				Text:   msg.Text,
				ReplyMarkup: types.InlineKeyboardMarkup{
					InlineKeyboard: [][]types.InlineKeyboardButton{
						buttons,
					},
				},
			})
		}
	}
}

func setScript(vm *goja.Runtime) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	content, err := os.ReadFile(dir + "\\examples\\js-bot\\handler.js")
	if err != nil {
		panic(err)
	}

	_, err = vm.RunString(string(content))
	if err != nil {
		panic(err)
	}
}
