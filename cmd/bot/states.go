package main

import (
	"context"
	"gote/internal/state"
	s "gote/internal/state"
	"gote/pkg/methods"
	"gote/pkg/types"
)

const (
	StartState = "Старт"
	GetName    = "Получить имя"
	SetName    = "Установить имя"
	GetMail    = "Получить почту"
	EndState   = "Конец"
	MenuState  = "Меню"
)

func createStateMachine() *s.StateMachine {
	startState := s.NewState(StartState, "/start", startAction)
	requestNameState := s.NewState(GetName, "/getName", requestName)
	setNameState := s.NewState(SetName, "/setName", setName)
	requestMailState := s.NewState(GetMail, "/getMail", requestMail)
	// endState := s.NewState(EndState, "/end")
	// menuState := s.NewState(MenuState, "/menu")

	startState.AddChildren(requestNameState)
	requestNameState.AddChildren(setNameState)
	setNameState.AddChildren(requestMailState)
	// requestMailState.AddChildren(endState)
	// endState.AddChildren(menuState)
	// menuState.AddChildren(startState)
	// menuState.AddChildren(endState)

	return s.NewStateMachine(startState, startState)
}

func startAction(ctx context.Context, update *types.Update, sm *state.StateMachine) {
	sm.NextState(ctx, update)
}

func requestName(ctx context.Context, update *types.Update, sm *state.StateMachine) {
	methods.SendMessage(ctx, types.SendMessage{
		ChatId: update.Message.Chat.Id,
		Text:   "Введите имя: ",
	})
}

func setName(ctx context.Context, update *types.Update, sm *state.StateMachine) {
	methods.SendMessage(ctx, types.SendMessage{
		ChatId: update.Message.Chat.Id,
		Text:   "Имя записано",
	})
	sm.NextState(ctx, update)
}

func requestMail(ctx context.Context, update *types.Update, sm *state.StateMachine) {
	methods.SendMessage(ctx, types.SendMessage{
		ChatId: update.Message.Chat.Id,
		Text:   "Введите почту: ",
	})
}
