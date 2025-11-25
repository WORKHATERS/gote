package state

import (
	"context"
	"gote/pkg/core"
	"gote/pkg/types"
)

type StateStore struct {
	States     *States
	UsersState *UsersState
}

type State struct {
	Name   string
	Handle Handler
}

type States map[string]*State
type UsersState map[int64]*State
type Handler func(context.Context, *types.Update, *core.BotContext)

func New() *StateStore {
	return &StateStore{
		States:     &States{},
		UsersState: &UsersState{},
	}
}

func (s *StateStore) Get(chatID int64) string {
	return ""
}

func (s *StateStore) Set(chatID int64, state string) {
}

func (s *StateStore) NewState(name string, handler Handler) *State {
	state := &State{
		Name:   name,
		Handle: handler,
	}
	(*s.States)[name] = state
	return state
}

func (s *StateStore) SetState(id int64, name string) {
	(*s.UsersState)[id] = (*s.States)[name]
}

func (s *StateStore) OnCommand(command string, state *State) {
	(*s.States)[command] = state
}
