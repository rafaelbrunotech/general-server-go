package user

import (
	"github.com/rafaelbrunoss/general-server-go/internal/common"
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/database"
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/application"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/repository"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/infrastructure/controller"
)

type UserRepository struct{}

type Container struct {
	UserRepository repository.IUserRepository
	UseCases       application.UseCases
	UserController controller.IUserController
}

func NewContainer(commonContainer *common.Container) *Container {
	logger := commonContainer.Services.Logger

	db := commonContainer.DB

	userRepository := provideRepositories(db)

	useCases := provideUseCases(logger, userRepository)

	controller := provideController(*useCases)

	return &Container{
		UserRepository: userRepository,
		UseCases:       *useCases,
		UserController: controller,
	}
}

func provideController(useCases application.UseCases) controller.IUserController {
	controller := controller.NewUserController(useCases)

	return controller
}

func provideRepositories(db *database.DB) repository.IUserRepository {
	repository := repository.NewUserRepository(db)

	return repository
}

func provideUseCases(logger service.ILogger, userRepository repository.IUserRepository) *application.UseCases {
	useCases := application.NewUseCases(logger, userRepository)

	return useCases
}
