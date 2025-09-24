package getuserbyid

import (
	"net/http"

	"github.com/rafaelbrunotech/general-server-go/internal/common/domain/model"
	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/service/logger"
	usererrors "github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/error"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/repository"
)

type GetUserByIdUseCase struct {
	logger         logger.ILogger
	userRepository repository.IUserRepository
}

func NewUseCase(
	logger logger.ILogger,
	userRepository repository.IUserRepository,
) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (u *GetUserByIdUseCase) Execute(request *GetUserByIdQuery) *model.ApiResponse[GetUserByIdResponse, string] {
	user, err := u.userRepository.GetUserById(request.UserId)

	if err != nil {
		return model.NewErrorApiResponse[GetUserByIdResponse, string]("user", usererrors.UserNotFound.Error(), http.StatusNotFound)
	}

	response, err := NewResponse(GetUserByIdResponseInput{User: *user})

	if err != nil {
		model.NewErrorApiResponse[GetUserByIdResponse, string]("response", err.Error(), http.StatusInternalServerError)
	}

	return model.NewSuccessApiResponse[GetUserByIdResponse, string](response, http.StatusOK)
}
