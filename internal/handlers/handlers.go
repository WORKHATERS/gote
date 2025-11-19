package handlers

import (
	"context"
	"gote/pkg/types"
)

type UpdateEvent int8

const (
	Message UpdateEvent = iota
	CallbackQuery
)

type HandlerFunc func(context.Context, types.Update)

type Handlers map[UpdateEvent]HandlerFunc

func NewHandlers() Handlers {
	return map[UpdateEvent]HandlerFunc{}
}

func (hs Handlers) Add(updateEvent UpdateEvent, handler HandlerFunc) {
	hs[updateEvent] = handler
}
