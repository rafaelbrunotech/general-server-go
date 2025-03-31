package getuserbyid

import (
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service/logger"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/repository"
)

type GetUserByIdUseCase struct {
	logger         logger.ILogger
	userRepository repository.IUserRepository
}

func NewGetUserByIdUseCase(
	logger logger.ILogger,
	userRepository repository.IUserRepository,
) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (u *GetUserByIdUseCase) Execute(request *GetUserByIdQuery) (*GetUserByIdResponse, error) {
	user, err := u.userRepository.GetUserById(request.UserId)

	if err != nil {
		return nil, err
	}

	response, err := NewGetUserByIdResponse(GetUserByIdResponseInput{User: *user})

	if err != nil {
		return nil, err
	}

	return response, nil
}
