package signup

type SignUpCommandInput struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignUpCommand struct {
	Email    string
	Name     string
	Password string
}

func NewSignUpCommand(input SignUpCommandInput) (*SignUpCommand, error) {
	return &SignUpCommand{
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
	}, nil
}
