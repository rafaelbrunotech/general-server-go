package signin

import (
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/entity"
)

type SignInResponseInput struct {
	AuthUser entity.AuthUser
}

type SignInResponse struct {
	AuthUser entity.AuthUser `json:"authUser"`
}

func NewSignInResponse(input SignInResponseInput) (*SignInResponse, error) {
	return &SignInResponse{AuthUser: input.AuthUser}, nil
}
