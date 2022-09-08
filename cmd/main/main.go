package main

import (
	"context"
	"gomoku/pkg/app"
	"gomoku/pkg/logger"
)

func main() {
	if err := app.Run(context.Background()); err != nil {
		logger.Error(err)
	}
}
