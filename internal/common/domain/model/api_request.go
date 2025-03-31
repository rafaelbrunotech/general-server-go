package model

import (
	"time"

	valueobject "github.com/rafaelbrunoss/general-server-go/internal/common/domain/value-object"
)

type ApiRequest struct {
	id        string
	AuthToken string
	Endpoint  string
	Timestamp time.Time
}

func NewApiRequest(
	authToken string,
	endpoint string,
	timestamp time.Time,
) *ApiRequest {
	return &ApiRequest{
		id:        valueobject.NewId().Value(),
		AuthToken: authToken,
		Endpoint:  endpoint,
		Timestamp: timestamp,
	}
}
