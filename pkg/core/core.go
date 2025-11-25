package core

import (
	"context"
	"gote/pkg/types"
)

type BotContext interface {
	Context() context.Context
	API() TelegramAPI
	Logger() Logger
	Router() Router
	State() MemoryState
	Store() MemoryStore
	Deps() Dependencies
}

type TelegramAPI interface {
	GetUpdates(ctx context.Context, param types.GetUpdates) ([]types.Update, error)
}

type Logger interface {
	Println(v ...any)
	Printf(format string, v ...any)
}

type Router interface {
	Process(ctx context.Context, b BotContext, upd types.Update)
}

type MemoryState interface {
	Get(chatID int64) string
	Set(chatID int64, state string)
}
type MemoryStore interface {
	Get(key string) any
	Set(key string, value any)
}

type Dependencies interface {
	Provide(obj any)
}
