package signin

import valueobject "github.com/rafaelbrunoss/general-server-go/internal/common/domain/value-object"

type SignInCommandInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInCommand struct {
	Email    *valueobject.Email
	Password string
}

func NewSignInCommand(input SignInCommandInput) (*SignInCommand, error) {
	email, err := valueobject.NewEmail(input.Email)

	if err != nil {
		return nil, err
	}

	return &SignInCommand{
		Email:    email,
		Password: input.Password,
	}, nil
}
