package getusers

import (
	"time"

	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/entity"
)

type userResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetUsersResponseInput struct {
	Users []entity.User
}

type GetUsersResponse struct {
	Users []userResponse `json:"users"`
}

func NewGetUsersResponse(input GetUsersResponseInput) (*GetUsersResponse, error) {
	var users []userResponse

	for _, user := range input.Users {
		userResponse := userResponse{
			Id:        user.Id.Value(),
			Email:     user.Email.Value(),
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		users = append(users, userResponse)
	}

	return &GetUsersResponse{Users: users}, nil
}
