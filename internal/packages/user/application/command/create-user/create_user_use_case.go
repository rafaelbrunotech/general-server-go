package createuser

import (
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/entity"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/repository"
)

type CreateUserUseCase struct {
	logger         service.ILogger
	userRepository repository.IUserRepository
}

func NewCreateUserUseCase(
	logger service.ILogger,
	userRepository repository.IUserRepository,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (u *CreateUserUseCase) Execute(request *CreateUserCommand) error {
	user, err := entity.NewUser(entity.UserInput{Name: request.Name})

	if err != nil {
		return err
	}

	err = u.userRepository.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}
