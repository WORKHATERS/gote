package state

import (
	"context"
	"gote/pkg/types"
)

type UsersState map[int64]*State
type States map[string]*State
type Action func(context.Context, *types.Update, *StateMachine)

type StateMachine struct {
	UsersState UsersState
	States     States
	StartState *State
	ResetState *State
}

type State struct {
	Name      string
	Condition string
	Action    Action
	parent    *State
	children  []*State
}

func NewState(name, condition string, action Action) *State {
	state := &State{
		Name:      name,
		Condition: condition,
		Action:    action,
	}

	return state
}

func (s *State) AddChildren(state *State) {
	s.children = append(s.children, state)
	state.parent = s
}

func NewStateMachine(start *State, reset *State) *StateMachine {
	return &StateMachine{
		UsersState: UsersState{},
		States:     States{},
		StartState: start,
		ResetState: reset,
	}
}

func (sm *StateMachine) NextState(ctx context.Context, update *types.Update) bool {
	id := update.Message.Chat.Id
	text := update.Message.Text
	state, ok := sm.UsersState[id]
	if !ok {
		state = sm.ResetState
		sm.UsersState[id] = sm.ResetState
	}

	children := state.children
	if len(children) == 0 {
		return false
	}

	if len(children) == 1 {
		sm.UsersState[id] = children[0]
		sm.UsersState[id].Action(ctx, update, sm)
		return true
	}

	for _, s := range children {
		if s.Condition == text {
			sm.UsersState[id] = s
			return true
		}
	}

	return false
}

func (sm *StateMachine) GetState(id int64) *State {
	return sm.UsersState[id]
}
