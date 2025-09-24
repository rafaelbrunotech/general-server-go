package updateuser

import (
	"net/http"

	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/service/logger"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/repository"
	model "github.com/rafaelbrunotech/general-server-go/internal/common/domain/model"
	usererrors "github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/error"
)

type UpdateUserUseCase struct {
	logger         logger.ILogger
	userRepository repository.IUserRepository
}

func NewUseCase(
	logger logger.ILogger,
	userRepository repository.IUserRepository,
) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (u *UpdateUserUseCase) Execute(request *UpdateUserCommand) *model.ApiResponse[any, string] {
	user, err := u.userRepository.GetUserById(request.UserId)

	if err != nil {
		return model.NewErrorApiResponse[any, string]("user", usererrors.UserNotFound.Error(), http.StatusNotFound)
	}

	user.SetName(request.Name)

	err = u.userRepository.UpdateUser(user)

	if err != nil {
		return model.NewErrorApiResponse[any, string]("user", err.Error(), http.StatusInternalServerError)
	}

	return model.NewSuccessApiResponse[any, string](nil, http.StatusOK)
}
