package bot

import (
	"context"
	"fmt"
	"gote/internal/commands"
	"gote/internal/handlers"
	"gote/internal/state"
	"gote/pkg/types"
	"log"
)

type Bot struct {
	Token        string
	ctx          context.Context
	offset       int64
	Commands     *commands.Commands
	Handlers     *handlers.Handlers
	StateMachine *state.StateMachine
}

func NewBot(token string) *Bot {
	bot := &Bot{
		Token: token,
		ctx:   context.Background(),
	}
	return bot
}

func (b *Bot) WithCommands(commands *commands.Commands) {
	b.Commands = commands
}

func (b *Bot) WithHandlers(handlers *handlers.Handlers) {
	b.Handlers = handlers
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
			response, err := bot.GetUpdates(bot.ctx, types.GetUpdates{
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
					id := update.Message.Chat.Id
					bot.StateMachine.SetState(&update)
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
