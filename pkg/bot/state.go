package bot

import (
	"context"
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
type Handler func(context.Context, *types.Update, *Bot)

func NewStateStore() *StateStore {
	return &StateStore{
		States:     &States{},
		UsersState: &UsersState{},
	}
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
