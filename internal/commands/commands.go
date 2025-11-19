package commands

import (
	"gote/internal/handlers"
)

type Commands map[string]handlers.HandlerFunc

func NewCommands() Commands {
	return map[string]handlers.HandlerFunc{}
}

func (c Commands) Add(command string, handler handlers.HandlerFunc) {
	c[command] = handler
}
