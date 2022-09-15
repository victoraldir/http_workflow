package app

import (
	"github.com/victoraldir/http-follower/internal/infra/config"
	"github.com/victoraldir/http-follower/internal/infra/terminal"
	"github.com/victoraldir/http-follower/internal/request/adapters"
	"github.com/victoraldir/http-follower/internal/request/usecases"
)

type ApplicationTerminal struct {
	CommandHandler terminal.CommandHandler
}

func NewApplicationTerminal(cfg config.Configuration) ApplicationTerminal {

	executeRequestFlowUseCase := Init(cfg)
	commandHandler := terminal.NewCommandHandler(executeRequestFlowUseCase)

	return ApplicationTerminal{
		CommandHandler: commandHandler,
	}
}

func Init(cfg config.Configuration) usecases.ExecuteRequestFlowUseCase {
	client := adapters.NewHTTPClient()
	return usecases.NewExecuteRequestFlowUseCase(client)
}
