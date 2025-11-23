package bot

import (
	"context"
	"gote/pkg/types"
)

type State struct {
	Name   string
	Handle Handler
}

type States map[string]*State
type UsersState map[int64]*State
type Handler func(context.Context, *types.Update, *Bot)

func (bot *Bot) OnState(name string, handler Handler) {
	(*bot.States)[name] = &State{
		Name:   name,
		Handle: handler,
	}
}

func (bot *Bot) SetState(id int64, name string) {
	(*bot.UsersState)[id] = (*bot.States)[name]
}
