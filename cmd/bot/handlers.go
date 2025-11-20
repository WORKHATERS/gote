package main

import (
	"context"
	"fmt"
	"gote/pkg/types"
)

func StartHandler(ctx context.Context, update types.Update) {
	fmt.Println("Я сказала стартуем!")
}

func MessageHandler(ctx context.Context, u types.Update) {
}
