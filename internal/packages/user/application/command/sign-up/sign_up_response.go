package signup

import (
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/entity"
)

type SignUpResponseInput struct {
	AuthUser entity.AuthUser
}

type SignUpResponse struct {
	AuthUser entity.AuthUser `json:"authUser"`
}

func NewResponse(input SignUpResponseInput) (*SignUpResponse, error) {
	return &SignUpResponse{AuthUser: input.AuthUser}, nil
}
