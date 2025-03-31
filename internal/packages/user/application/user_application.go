package application

import (
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service/logger"
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service/tokenizer"
	signin "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/command/sign-in"
	signup "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/command/sign-up"
	updateuser "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/command/update-user"
	getuserbyid "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/query/get-user-by-id"
	getusers "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/query/get-users"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/repository"
)

type CommandUseCases struct {
	SignIn     *signin.SignInUseCase
	SignUp     *signup.SignUpUseCase
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

func newCommandUseCases(logger logger.ILogger, tokenizer tokenizer.ITokenizer, userRepository repository.IUserRepository) *CommandUseCases {
	signIn := signin.NewSignInUseCase(logger, tokenizer, userRepository)
	signUp := signup.NewSignUpUseCase(logger, tokenizer, userRepository)
	updateUser := updateuser.NewUpdateUserUseCase(logger, userRepository)

	return &CommandUseCases{
		SignIn:     signIn,
		SignUp:     signUp,
		UpdateUser: updateUser,
	}
}

func newQueryUseCases(logger logger.ILogger, userRepository repository.IUserRepository) *QueryUseCases {
	getUserById := getuserbyid.NewGetUserByIdUseCase(logger, userRepository)
	getUsers := getusers.NewGetUsersUseCase(logger, userRepository)

	return &QueryUseCases{
		GetUserById: getUserById,
		GetUsers:    getUsers,
	}
}

func NewUseCases(logger logger.ILogger, tokenizer tokenizer.ITokenizer, userRepository repository.IUserRepository) *UseCases {
	command := newCommandUseCases(logger, tokenizer, userRepository)

	query := newQueryUseCases(logger, userRepository)

	return &UseCases{
		Command: command,
		Query:   query,
	}
}
