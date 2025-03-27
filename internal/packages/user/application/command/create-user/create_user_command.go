package createuser

type CreateUserCommandInput struct {
	Name string `json:"name"`
}

type CreateUserCommand struct {
	Name string
}

func NewCreateUserCommand(input CreateUserCommandInput) (*CreateUserCommand, error) {
	return &CreateUserCommand{
		Name: input.Name,
	}, nil
}
