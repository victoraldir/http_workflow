package domain

type Request struct {
	Method      string
	URL         string
	Body        string
	Headers     map[string]string
	RetryPeriod int
	Assertion   Assertion
}
