package bot

import (
	"context"
	"fmt"
	"gote/internal/commands"
	"gote/internal/handlers"
	"gote/internal/state"
	"gote/pkg/methods"
	"gote/pkg/types"
	"log"
)

type Bot struct {
	ctx          context.Context
	offset       int64
	Commands     *commands.Commands
	Handlers     *handlers.Handlers
	StateMachine *state.StateMachine
}

func NewBot(ctx context.Context) *Bot {
	bot := &Bot{
		ctx:      ctx,
		Commands: &commands.Commands{},
	}
	return bot
}

func (bot *Bot) OnCommand(command string, handler handlers.HandlerFunc) {
	(*bot.Commands)[command] = handler
}

func (b *Bot) WithState(stateMachine *state.StateMachine) {
	b.StateMachine = stateMachine
}

func (bot *Bot) RunUpdate() {
	for {
		select {
		case <-bot.ctx.Done():
			return
		default:
			response, err := methods.GetUpdates(bot.ctx, types.GetUpdates{
				Limit:   100,
				Timeout: 50,
				Offset:  bot.offset,
			})
			if err != nil {
				log.Println("Не получилось получить Update")
				return
			}

			go func(ctx context.Context, updates []types.Update) {
				for _, update := range updates {
					msg := update.Message
					if msg == nil {
						continue
					}
					id := update.Message.Chat.Id

					text := msg.Text
					handlerFunc, ok := (*bot.Commands)[text]
					if ok {
						// bot.StateMachine.UsersState[id]
						// bot.StateMachine.UsersState[id].Action(ctx, &update)
						handlerFunc(ctx, update)
						continue
					}

					if bot.StateMachine != nil {
						bot.StateMachine.NextState(ctx, &update)
					}
					fmt.Println(bot.StateMachine.GetState(id))
				}
			}(bot.ctx, response)

			lenUpdate := len(response)
			if lenUpdate > 0 {
				bot.offset = response[lenUpdate-1].UpdateId + 1
			}
		}
	}
}
