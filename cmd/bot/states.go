package main

import s "gote/internal/state"

const (
	StartState = "Старт"
	EndState   = "Конец"
	MenuState  = "Меню"
)

func createStateMachine() *s.StateMachine {
	startState := s.NewState(StartState, "/start")
	endState := s.NewState(EndState, "/end")
	menuState := s.NewState(MenuState, "/menu")

	startState.AddChildren(endState)
	endState.AddChildren(menuState)
	menuState.AddChildren(startState)
	menuState.AddChildren(endState)

	return &s.StateMachine{
		States:     s.States{},
		StartState: startState,
		ResetState: menuState,
	}
}
