package updateuser

import (
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/repository"
)

type UpdateUserUseCase struct {
	logger         service.ILogger
	userRepository repository.IUserRepository
}

func NewUpdateUserUseCase(
	logger service.ILogger,
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

	user.UpdateName(request.Name)

	err = u.userRepository.UpdateUser(user)

	if err != nil {
		return err
	}

	return nil
}
