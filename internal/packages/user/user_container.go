package user

import (
	"github.com/rafaelbrunotech/general-server-go/internal/common"
	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/database"
	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/service/logger"
	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/service/tokenizer"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/application"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/repository"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/infrastructure/controller"
)

type UserRepository struct{}

type Container struct {
	UserRepository repository.IUserRepository
	UseCases       application.UseCases
	AuthController controller.IAuthController
	UserController controller.IUserController
}

func NewContainer(commonContainer *common.Container) *Container {
	logger := commonContainer.Services.Logger
	tokenizer := commonContainer.Services.Tokenizer

	db := commonContainer.DB

	userRepository := provideRepositories(db)

	useCases := provideUseCases(logger, tokenizer, userRepository)

	authController := provideAuthController(*useCases)
	userController := provideUserController(*useCases)

	return &Container{
		UserRepository: userRepository,
		UseCases:       *useCases,
		AuthController: authController,
		UserController: userController,
	}
}

func provideAuthController(useCases application.UseCases) controller.IAuthController {
	controller := controller.NewAuthController(useCases)

	return controller
}

func provideUserController(useCases application.UseCases) controller.IUserController {
	controller := controller.NewUserController(useCases)

	return controller
}

func provideRepositories(db *database.DB) repository.IUserRepository {
	repository := repository.NewUserRepository(db)

	return repository
}

func provideUseCases(logger logger.ILogger, tokenizer tokenizer.ITokenizer, userRepository repository.IUserRepository) *application.UseCases {
	useCases := application.NewUseCases(logger, tokenizer, userRepository)

	return useCases
}
