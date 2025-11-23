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
	WorkersCount   int64
}

type Bot struct {
	ctx          context.Context
	API          *api.API
	State        *StateStore
	Store        *Store
	Dependencies *Dependencies
	updateParams types.GetUpdates
	workers      chan struct{}
}

func NewBot(ctx context.Context, config Config) *Bot {
	if config.Limit <= 0 {
		config.Limit = 100
	}
	if config.Timeout <= 0 {
		config.Timeout = 50
	}
	if config.WorkersCount <= 0 {
		config.WorkersCount = 100
	}

	return &Bot{
		ctx:     ctx,
		API:     api.NewAPI(config.Token),
		State:   NewStateStore(),
		Store:   NewStore(),
		workers: make(chan struct{}, config.WorkersCount),
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
		}

		updates, err := bot.API.GetUpdates(bot.ctx, bot.updateParams)
		if err != nil {
			return err
		}

		go func(uu []types.Update) {
			for _, u := range uu {
				bot.workers <- struct{}{}

				go func(update types.Update) {
					defer func() { <-bot.workers }()
					handleUpdate(bot.ctx, update, bot)
				}(u)
			}
		}(updates)

		lenUpdate := len(updates)
		if lenUpdate > 0 {
			bot.updateParams.Offset = updates[lenUpdate-1].UpdateId + 1
		}
	}
}

func handleUpdate(ctx context.Context, update types.Update, bot *Bot) {
	id, isFound := getChatID(update)
	if !isFound {
		return
	}

	if msg := update.Message; msg != nil {
		text := msg.Text

		state, ok := (*bot.State.States)[text]
		if ok {
			(*bot.State.UsersState)[id] = state
			state.Handle(ctx, &update, bot)
			return
		}
	}

	userState := (*bot.State.UsersState)[id]
	if userState != nil {
		userState.Handle(ctx, &update, bot)
	}
}

func getChatID(u types.Update) (int64, bool) {
	var id int64
	founded := true

	switch {
	case u.Message != nil:
		id = u.Message.Chat.Id

	case u.EditedMessage != nil:
		id = u.EditedMessage.Chat.Id

	case u.ChannelPost != nil:
		id = u.ChannelPost.Chat.Id

	case u.EditedChannelPost != nil:
		id = u.EditedChannelPost.Chat.Id

	case u.BusinessMessage != nil:
		id = u.BusinessMessage.Chat.Id

	case u.EditedBusinessMessage != nil:
		id = u.EditedBusinessMessage.Chat.Id

	case u.DeletedBusinessMessages != nil:
		id = u.DeletedBusinessMessages.Chat.Id

	case u.CallbackQuery != nil && u.CallbackQuery.Message != nil:
		msg := (*u.CallbackQuery.Message)
		switch m := msg.(type) {
		case *types.Message:
			id = m.Chat.Id
		case *types.InaccessibleMessage:
			id = m.Chat.Id
		}

	case u.MyChatMember != nil:
		id = u.MyChatMember.Chat.Id

	case u.ChatMember != nil:
		id = u.ChatMember.Chat.Id

	case u.ChatJoinRequest != nil:
		id = u.ChatJoinRequest.Chat.Id

	case u.MessageReaction != nil:
		id = u.MessageReaction.Chat.Id

	case u.MessageReactionCount != nil:
		id = u.MessageReactionCount.Chat.Id

	case u.ChatBoost != nil:
		id = u.ChatBoost.Chat.Id

	case u.RemovedChatBoost != nil:
		id = u.RemovedChatBoost.Chat.Id

	default:
		founded = false
	}

	return id, founded
}
