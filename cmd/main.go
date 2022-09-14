package main

import (
	"fmt"
	"os"

	"github.com/victoraldir/http-follower/internal/request/adapters"
	"github.com/victoraldir/http-follower/internal/request/usecases"
	"github.com/victoraldir/http-follower/pkg/file"
)

func main() {

	arg := os.Args[1]

	workflowRequest, err := file.LoadWorkflowrequestFromYaml(arg)

	if err != nil {
		fmt.Println(err)
	}

	client := adapters.NewHTTPClient()
	executor := usecases.NewExecuteRequestFlowUseCase(client)

	executor.Execute(*workflowRequest)
}
