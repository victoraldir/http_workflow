package domain

import (
	"fmt"
)

type Assertion struct {
	Name               string `json:"name"`
	ExpectedCode       int    `json:"expected"`
	OnFailure          string `json:"onfailure"`
	MinValidAssertions int    `json:"minValidAssertions"`
}

// Validate validates the assertion
func (a *Assertion) Validate(response *Response) error {
	if a.ExpectedCode != response.StatusCode {
		return fmt.Errorf("Assertion failed. Expected: %d Got: %d", a.ExpectedCode, response.StatusCode)
	}

	return nil
}
