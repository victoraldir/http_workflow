package usecases

import (
	"fmt"
	"math"
	"time"

	"github.com/victoraldir/http-follower/internal/request/core/ports"
	"github.com/victoraldir/http-follower/internal/request/dto"
	"github.com/victoraldir/http-follower/pkg/logger"
)

type (
	ExecuteRequestFlowUseCase interface {
		Execute(workflowRequest dto.WorkflowRequest) error
	}

	executeRequestFlowUseCase struct {
		client ports.Client
	}
)

func NewExecuteRequestFlowUseCase(client ports.Client) ExecuteRequestFlowUseCase {
	return &executeRequestFlowUseCase{
		client: client,
	}
}

func (e *executeRequestFlowUseCase) Execute(workflowRequest dto.WorkflowRequest) error {

	logger.Info("[Executing workflow...]")

	for _, requestPlan := range workflowRequest.Requests {
		request := requestPlan.ToRequest()
		assertion := request.Assertion
		minValidAssertions := int(math.Max(float64(assertion.MinValidAssertions), 1))
		validAssertionCount := 0

		logger.Info(fmt.Sprintf("[Executing request] %s \n", requestPlan.Request))

		for {

			resp, err := e.client.Do(&request)
			logger.Debug(fmt.Sprintf("[Response] %v \n", resp))
			time.Sleep(time.Duration(request.RetryPeriod) * time.Second)

			if err != nil {

				if assertion.OnFailure == "retry" {
					logger.Info(fmt.Sprintf("[Request failed] Got error: %s. Retrying request: %s \n", err.Error(), requestPlan.Request))
					validAssertionCount = 0
					continue
				}

				if assertion.OnFailure == "skip" {
					logger.Info(fmt.Sprintf("Assertion failed: %s. Skipping... \n", err.Error()))
					break
				}

				logger.Error("Error executing request", err)
				return err
			}

			logger.Info(fmt.Sprintf("[Executing assertion] %s response code: %d", assertion.Name, resp.StatusCode))
			err = assertion.Validate(resp)

			if err != nil {

				if assertion.OnFailure == "retry" {
					logger.Info(fmt.Sprintf("[Assertion failed] %s. Retrying... \n", err.Error()))
					validAssertionCount = 0
					continue
				}

				if assertion.OnFailure == "skip" {
					logger.Info(fmt.Sprintf("Assertion failed: %s. Skipping... \n", err.Error()))
					break
				}

				logger.Info(fmt.Sprintf("Assertion failed: %s. Breaking execution... \n", err.Error()))
				return err
			}

			validAssertionCount++

			logger.Info(
				fmt.Sprintf("Assertion %s passed [%d/%d] \n",
					assertion.Name, validAssertionCount, minValidAssertions,
				),
			)

			if validAssertionCount >= minValidAssertions {
				break
			}
		}
	}

	return nil
}
