package state

import "gote/pkg/types"

type States map[int64]*State

type StateMachine struct {
	States     States
	StartState *State
	ResetState *State
}

type State struct {
	Name      string
	Condition string
	parent    *State
	children  []*State
}

func NewState(name, condition string) *State {
	return &State{
		Name:      name,
		Condition: condition,
	}
}

func (s *State) AddChildren(state *State) {
	s.children = append(s.children, state)
	state.parent = s
}

func NewStateMachine(start *State, reset *State) *StateMachine {
	return &StateMachine{
		States:     States{},
		StartState: start,
		ResetState: reset,
	}
}

func (sm *StateMachine) SetState(update *types.Update) bool {
	id := update.Message.Chat.Id
	text := update.Message.Text
	state, ok := sm.States[id]
	if !ok {
		sm.States[id] = sm.ResetState
		return false
	}

	children := state.children
	if len(children) == 0 {
		return false
	}

	if len(children) == 1 {
		sm.States[id] = children[0]
		return true
	}

	for _, s := range children {
		if s.Condition == text {
			sm.States[id] = s
			return true
		}
	}

	return false
}

func (sm *StateMachine) GetState(id int64) *State {
	return sm.States[id]
}
