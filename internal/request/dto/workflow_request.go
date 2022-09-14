package dto

import "github.com/victoraldir/http-follower/internal/request/core/domain"

const (
	// Default RetryPeriod in seconds
	DefaultRetryPeriod = 1
)

type WorkflowRequest struct {
	Requests []RequestPlan `json:"workflow" yaml:"workflow"`
}

type RequestPlan struct {
	Request        string         `json:"request" yaml:"request"`
	Method         string         `json:"method" yaml:"method"`
	Url            string         `json:"url" yaml:"url"`
	Body           string         `json:"body" yaml:"body"`
	RetryPeriod    int            `json:"retry_period" yaml:"retry_period"`
	AsssertionPlan AsssertionPlan `json:"assertion" yaml:"assertion"`
}

type AsssertionPlan struct {
	Name         string `json:"name" yaml:"name"`
	ExpectedCode int    `json:"expectedcode" yaml:"expectedcode"`
	OnFailure    string `json:"onfailure" yaml:"onfailure"`
}

// Create function that converts RequestPlan to Request
func (r *RequestPlan) ToRequest() domain.Request {

	if r.RetryPeriod == 0 {
		r.RetryPeriod = DefaultRetryPeriod
	}

	return domain.Request{
		Method:      r.Method,
		URL:         r.Url,
		Body:        r.Body,
		RetryPeriod: r.RetryPeriod,
		Assertion:   r.AsssertionPlan.ToAsssertion(),
	}
}

// create function that converts AssertionPlan to Assertion
func (a *AsssertionPlan) ToAsssertion() domain.Assertion {
	return domain.Assertion{
		Name:         a.Name,
		ExpectedCode: a.ExpectedCode,
		OnFailure:    a.OnFailure,
	}
}
