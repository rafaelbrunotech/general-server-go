package getuserbyid

import valueobject "github.com/rafaelbrunoss/general-server-go/internal/common/domain/value-object"

type GetUserByIdQueryInput struct {
	UserId string
}

type GetUserByIdQuery struct {
	UserId *valueobject.Id
}

func NewGetUserByIdQuery(input GetUserByIdQueryInput) (*GetUserByIdQuery, error) {
	return &GetUserByIdQuery{
		UserId: valueobject.NewValue(input.UserId),
	}, nil
}
