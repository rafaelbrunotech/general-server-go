package application

import (
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service"
	createuser "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/command/create-user"
	updateuser "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/command/update-user"
	getuserbyid "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/query/get-user-by-id"
	getusers "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/query/get-users"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/repository"
)

type CommandUseCases struct {
	CreateUser *createuser.CreateUserUseCase
	UpdateUser *updateuser.UpdateUserUseCase
}

type QueryUseCases struct {
	GetUserById *getuserbyid.GetUserByIdUseCase
	GetUsers    *getusers.GetUsersUseCase
}

type UseCases struct {
	Command *CommandUseCases
	Query   *QueryUseCases
}

func newCommandUseCases(logger service.ILogger, userRepository repository.IUserRepository) *CommandUseCases {
	createUser := createuser.NewCreateUserUseCase(logger, userRepository)
	updateUser := updateuser.NewUpdateUserUseCase(logger, userRepository)

	return &CommandUseCases{
		CreateUser: createUser,
		UpdateUser: updateUser,
	}
}

func newQueryUseCases(logger service.ILogger, userRepository repository.IUserRepository) *QueryUseCases {
	getUserById := getuserbyid.NewGetUserByIdUseCase(logger, userRepository)
	getUsers := getusers.NewGetUsersUseCase(logger, userRepository)

	return &QueryUseCases{
		GetUserById: getUserById,
		GetUsers:    getUsers,
	}
}

func NewUseCases(logger service.ILogger, userRepository repository.IUserRepository) *UseCases {
	command := newCommandUseCases(logger, userRepository)

	query := newQueryUseCases(logger, userRepository)

	return &UseCases{
		Command: command,
		Query:   query,
	}
}
