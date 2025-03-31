package updateuser

import (
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service/logger"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/repository"
)

type UpdateUserUseCase struct {
	logger         logger.ILogger
	userRepository repository.IUserRepository
}

func NewUpdateUserUseCase(
	logger logger.ILogger,
	userRepository repository.IUserRepository,
) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (u *UpdateUserUseCase) Execute(request *UpdateUserCommand) error {
	user, err := u.userRepository.GetUserById(request.UserId)

	if err != nil {
		return err
	}

	user.SetName(request.Name)

	err = u.userRepository.UpdateUser(user)

	if err != nil {
		return err
	}

	return nil
}
