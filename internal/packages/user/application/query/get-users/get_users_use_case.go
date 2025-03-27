package getusers

import (
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/repository"
)

type GetUsersUseCase struct {
	logger         service.ILogger
	userRepository repository.IUserRepository
}

func NewGetUsersUseCase(
	logger service.ILogger,
	userRepository repository.IUserRepository,
) *GetUsersUseCase {
	return &GetUsersUseCase{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (u *GetUsersUseCase) Execute(request *GetUsersQuery) (*GetUsersResponse, error) {
	users, err := u.userRepository.GetUsers()

	if err != nil {
		return nil, err
	}

	response, err := NewGetUsersResponse(GetUsersResponseInput{Users: users})

	if err != nil {
		return nil, err
	}

	return response, nil
}
