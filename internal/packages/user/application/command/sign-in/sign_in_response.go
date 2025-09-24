package signin

import (
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/entity"
)

type SignInResponseInput struct {
	AuthUser entity.AuthUser
}

type SignInResponse struct {
	AuthUser entity.AuthUser `json:"authUser"`
}

func NewResponse(input SignInResponseInput) (*SignInResponse, error) {
	return &SignInResponse{AuthUser: input.AuthUser}, nil
}
