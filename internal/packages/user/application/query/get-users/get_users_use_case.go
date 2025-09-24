package getusers

import (
	"net/http"

	"github.com/rafaelbrunotech/general-server-go/internal/common/domain/model"
	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/service/logger"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/repository"
	usererrors "github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/error"
)

type GetUsersUseCase struct {
	logger         logger.ILogger
	userRepository repository.IUserRepository
}

func NewUseCase(
	logger logger.ILogger,
	userRepository repository.IUserRepository,
) *GetUsersUseCase {
	return &GetUsersUseCase{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (u *GetUsersUseCase) Execute(request *GetUsersQuery) *model.ApiResponse[GetUsersResponse, string] {
	users, err := u.userRepository.GetUsers()

	if err != nil {
		return model.NewErrorApiResponse[GetUsersResponse, string]("user", usererrors.UserNotFound.Error(), http.StatusNotFound)
	}

	response, err := NewResponse(GetUsersResponseInput{Users: users})

	if err != nil {
		return model.NewErrorApiResponse[GetUsersResponse, string]("response", err.Error(), http.StatusInternalServerError)
	}

	return model.NewSuccessApiResponse[GetUsersResponse, string](response, http.StatusOK)
}
