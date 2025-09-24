package signup

import (
	"net/http"

	"github.com/rafaelbrunotech/general-server-go/internal/common/domain/model"
	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/service/logger"
	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/service/tokenizer"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/entity"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/repository"
)

type SignUpUseCase struct {
	logger         logger.ILogger
	tokenizer      tokenizer.ITokenizer
	userRepository repository.IUserRepository
}

func NewUseCase(
	logger logger.ILogger,
	tokenizer tokenizer.ITokenizer,
	userRepository repository.IUserRepository,
) *SignUpUseCase {
	return &SignUpUseCase{
		logger:         logger,
		tokenizer:      tokenizer,
		userRepository: userRepository,
	}
}

func (u *SignUpUseCase) Execute(request *SignUpCommand) *model.ApiResponse[SignUpResponse, string] {
	user, err := entity.NewUser(
		entity.UserInput{
			Email:    request.Email,
			Name:     request.Name,
			Password: request.Password,
		},
	)

	if err != nil {
		model.NewErrorApiResponse[SignUpResponse, string]("user", err.Error(), http.StatusUnprocessableEntity)
	}

	err = u.userRepository.CreateUser(user)

	if err != nil {
		return model.NewErrorApiResponse[SignUpResponse, string]("user", err.Error(), http.StatusInternalServerError)
	}

	tokenData := *model.NewTokenData(user.Id.Value(), user.Email.Value())
	accessToken, err := u.tokenizer.GenerateAccessToken(tokenData)

	if err != nil {
		return model.NewErrorApiResponse[SignUpResponse, string]("refreshToken", err.Error(), http.StatusInternalServerError)
	}

	refreshToken, err := u.tokenizer.GenerateRefreshToken(tokenData)

	if err != nil {
		return model.NewErrorApiResponse[SignUpResponse, string]("accessToken", err.Error(), http.StatusInternalServerError)
	}

	authUser := entity.NewAuthUser(entity.AuthUserInput{
		Id:           user.Id,
		Name:         user.Name,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

	response, err := NewResponse(SignUpResponseInput{AuthUser: authUser})

	if err != nil {
		return model.NewErrorApiResponse[SignUpResponse, string]("response", err.Error(), http.StatusInternalServerError)
	}

	return model.NewSuccessApiResponse[SignUpResponse, string](response, http.StatusCreated)
}
