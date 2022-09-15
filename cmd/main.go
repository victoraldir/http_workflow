package main

import (
	"github.com/victoraldir/http-follower/app"
	"github.com/victoraldir/http-follower/internal/infra/config"
	"github.com/victoraldir/http-follower/pkg/logger"
)

func main() {

	cfg := config.InitConfiguration()
	logger.Init(cfg.LogLevel)
	app := app.NewApplicationTerminal(cfg)

	commandHandler := app.CommandHandler
	commandHandler.Handle()
}
