package updateuser

import valueobject "github.com/rafaelbrunoss/general-server-go/internal/common/domain/value-object"

type UpdateUserCommandInput struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

type UpdateUserCommand struct {
	UserId *valueobject.Id
	Name   string
}

func NewUpdateUserCommand(input UpdateUserCommandInput) (*UpdateUserCommand, error) {
	return &UpdateUserCommand{
		UserId: valueobject.NewValue(input.UserId),
		Name:   input.Name,
	}, nil
}
