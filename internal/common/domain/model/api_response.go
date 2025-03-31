package model

import "os"

type ApiResponseMetadata struct {
	Environment string `json:"environment"`
	Page        uint32 `json:"page"`
	PerPage     uint32 `json:"perPage"`
	Total       uint32 `json:"total"`
}

type ApiResponseError[E any] struct {
	Details E      `json:"details"`
	Message string `json:"message"`
}

type ApiResponse[T any, E any] struct {
	Data     T                    `json:"data"`
	Error    *ApiResponseError[E] `json:"error"`
	Metadata ApiResponseMetadata  `json:"metadata"`
	Status   uint16               `json:"status"`
}

func NewSuccessApiResponse[T any](
	data T,
	status uint16,
) *ApiResponse[T, any] {
	r := &ApiResponse[T, any]{
		Data:  data,
		Error: nil,
		Metadata: ApiResponseMetadata{
			Environment: os.Getenv("ENV"),
		},
		Status: status,
	}

	return r
}

func NewErrorApiResponse[E any](
	errorDetails E,
	message string,
	status uint16,
) *ApiResponse[any, E] {
	r := &ApiResponse[any, E]{
		Data: nil,
		Error: &ApiResponseError[E]{
			Message: message,
			Details: errorDetails,
		},
		Metadata: ApiResponseMetadata{
			Environment: os.Getenv("ENV"),
		},
		Status: status,
	}

	return r
}
