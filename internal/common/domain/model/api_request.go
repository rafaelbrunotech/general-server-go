package model

import (
	"time"

	valueobject "github.com/rafaelbrunotech/general-server-go/internal/common/domain/value-object"
)

type ApiRequest struct {
	id        string
	AuthToken string
	Timestamp time.Time
}

func NewApiRequest(
	authToken string,
	timestamp time.Time,
) *ApiRequest {
	return &ApiRequest{
		id:        valueobject.NewId().Value(),
		AuthToken: authToken,
		Timestamp: timestamp,
	}
}
