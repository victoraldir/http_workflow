package domain

type Request struct {
	Method      string
	URL         string
	Body        string
	RetryPeriod int
	Assertion   Assertion
}
