package application

import (
	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/service/logger"
	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/service/tokenizer"
	signin "github.com/rafaelbrunotech/general-server-go/internal/packages/user/application/command/sign-in"
	signup "github.com/rafaelbrunotech/general-server-go/internal/packages/user/application/command/sign-up"
	updateuser "github.com/rafaelbrunotech/general-server-go/internal/packages/user/application/command/update-user"
	getuserbyid "github.com/rafaelbrunotech/general-server-go/internal/packages/user/application/query/get-user-by-id"
	getusers "github.com/rafaelbrunotech/general-server-go/internal/packages/user/application/query/get-users"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/repository"
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
	signIn := signin.NewUseCase(logger, tokenizer, userRepository)
	signUp := signup.NewUseCase(logger, tokenizer, userRepository)
	updateUser := updateuser.NewUseCase(logger, userRepository)

	return &CommandUseCases{
		SignIn:     signIn,
		SignUp:     signUp,
		UpdateUser: updateUser,
	}
}

func newQueryUseCases(logger logger.ILogger, userRepository repository.IUserRepository) *QueryUseCases {
	getUserById := getuserbyid.NewUseCase(logger, userRepository)
	getUsers := getusers.NewUseCase(logger, userRepository)

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
