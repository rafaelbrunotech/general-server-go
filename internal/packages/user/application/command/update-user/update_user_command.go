package updateuser

import valueobject "github.com/rafaelbrunotech/general-server-go/internal/common/domain/value-object"

type UpdateUserCommandInput struct {
	UserId   string `json:"userId"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UpdateUserCommand struct {
	UserId   *valueobject.Id
	Email    string
	Name     string
	Password string
}

func NewCommand(input UpdateUserCommandInput) (*UpdateUserCommand, error) {
	return &UpdateUserCommand{
		UserId:   valueobject.NewValue(input.UserId),
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
	}, nil
}
