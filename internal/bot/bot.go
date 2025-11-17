package bot

import (
	"context"
	"gote/pkg/types"
	"log"
)

type Bot struct {
	Token  string
	ctx    context.Context
	offset int64
}

func NewBot(token string) *Bot {
	return &Bot{
		Token: token,
		ctx:   context.Background(),
	}
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
				log.Println("Чота пошло не так")
				return
			}

			go func(ctx context.Context, updates []types.Update) {
				for _, update := range updates {
					log.Println(update.Message.Text)
				}
			}(bot.ctx, response)

			lenUpdate := len(response)
			if lenUpdate > 0 {
				bot.offset = response[lenUpdate-1].UpdateId + 1
			}
		}
	}
}
