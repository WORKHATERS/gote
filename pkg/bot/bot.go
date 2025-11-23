package bot

import (
	"context"
	"gote/pkg/api"
	"gote/pkg/types"
)

type Config struct {
	Token          string
	Limit          int64
	Timeout        int64
	Offset         int64
	AllowedUpdates []string
}

type Bot struct {
	ctx          context.Context
	updateParams types.GetUpdates
	API          *api.API
	State        *StateStore
	Store        *Store
	Dependencies *Dependencies
}

func NewBot(ctx context.Context, config Config) *Bot {
	if config.Limit <= 0 {
		config.Limit = 100
	}
	if config.Timeout <= 0 {
		config.Timeout = 50
	}
	return &Bot{
		ctx:   ctx,
		API:   api.NewAPI(config.Token),
		State: NewStateStore(),
		Store: NewStore(),
		updateParams: types.GetUpdates{
			Limit:          config.Limit,
			Timeout:        config.Timeout,
			Offset:         config.Offset,
			AllowedUpdates: config.AllowedUpdates,
		},
	}
}

func (bot *Bot) AddDependencies(dd *Dependencies) {
	bot.Dependencies = dd
}

func (bot *Bot) Run() error {
	for {
		select {
		case <-bot.ctx.Done():
			return bot.ctx.Err()
		default:
			response, err := bot.API.GetUpdates(bot.ctx, bot.updateParams)
			if err != nil {
				return err
			}

			go func(ctx context.Context, updates []types.Update) {
				for _, update := range updates {
					msg := update.Message
					if msg == nil {
						continue
					}

					id := msg.Chat.Id
					text := msg.Text

					state, ok := (*bot.State.States)[text]
					if ok {
						(*bot.State.UsersState)[id] = state
						state.Handle(ctx, &update, bot)
						continue
					}

					userState := (*bot.State.UsersState)[id]
					if userState != nil {
						userState.Handle(ctx, &update, bot)
					}
				}
			}(bot.ctx, response)

			lenUpdate := len(response)
			if lenUpdate > 0 {
				bot.updateParams.Offset = response[lenUpdate-1].UpdateId + 1
			}
		}
	}
}
