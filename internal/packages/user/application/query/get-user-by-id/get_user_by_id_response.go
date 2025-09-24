package getuserbyid

import (
	"time"

	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/entity"
)

type userResponse struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetUserByIdResponseInput struct {
	User entity.User
}

type GetUserByIdResponse struct {
	User userResponse `json:"user"`
}

func NewResponse(input GetUserByIdResponseInput) (*GetUserByIdResponse, error) {
	userResponse := userResponse{
		Id:        input.User.Id.Value(),
		Email:     input.User.Email.Value(),
		Name:      input.User.Name,
		CreatedAt: input.User.CreatedAt,
		UpdatedAt: input.User.UpdatedAt,
	}

	return &GetUserByIdResponse{User: userResponse}, nil
}
