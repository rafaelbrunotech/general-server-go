package model

import "time"

type ApiResult struct {
	Data        any
	Endpoint    string
	Environment string
	Messages    []string
	Status      uint16
	Success     bool
	Timestamp   time.Time
}

func NewApiResult(
	data any,
	endpoint string,
	environment string,
	messages []string,
	status uint16,
	success bool,
) *ApiResult {
	r := &ApiResult{
		Data:        data,
		Endpoint:    endpoint,
		Environment: environment,
		Messages:    messages,
		Status:      status,
		Success:     success,
		Timestamp:   time.Now(),
	}

	return r
}
