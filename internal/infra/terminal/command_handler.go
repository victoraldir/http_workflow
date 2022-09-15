package terminal

import (
	"io/ioutil"
	"os"

	"github.com/victoraldir/http-follower/internal/request/dto"
	"github.com/victoraldir/http-follower/internal/request/usecases"
	"github.com/victoraldir/http-follower/pkg/logger"
	"gopkg.in/yaml.v2"
)

type CommandHandler struct {
	FlowExecutor usecases.ExecuteRequestFlowUseCase
}

func NewCommandHandler(flowExecutor usecases.ExecuteRequestFlowUseCase) CommandHandler {
	return CommandHandler{
		FlowExecutor: flowExecutor,
	}
}

func (c CommandHandler) Handle() {
	bytes := ReadStdin()
	workflowRequest, err := ByteToWorkflow(bytes)

	if err != nil {
		logger.Panic("Error parsing workflow file: ", err)
	}

	c.FlowExecutor.Execute(*workflowRequest)
}

func ReadStdin() []byte {
	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		logger.Error("Error reading workflow file: ", err)
	}

	return bytes
}

func ByteToWorkflow(fileContent []byte) (*dto.WorkflowRequest, error) {

	var workflow *dto.WorkflowRequest

	err := yaml.Unmarshal(fileContent, &workflow)
	if err != nil {
		return nil, err
	}

	return workflow, nil
}
