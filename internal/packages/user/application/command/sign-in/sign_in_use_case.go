package signin

import (
	"net/http"

	"github.com/rafaelbrunotech/general-server-go/internal/common/domain/model"
	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/service/logger"
	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/service/tokenizer"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/entity"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/repository"
	usererrors "github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/error"
)

type SignInUseCase struct {
	logger         logger.ILogger
	tokenizer      tokenizer.ITokenizer
	userRepository repository.IUserRepository
}

func NewUseCase(
	logger logger.ILogger,
	tokenizer tokenizer.ITokenizer,
	userRepository repository.IUserRepository,
) *SignInUseCase {
	return &SignInUseCase{
		logger:         logger,
		tokenizer:      tokenizer,
		userRepository: userRepository,
	}
}

func (u *SignInUseCase) Execute(request *SignInCommand) *model.ApiResponse[SignInResponse, string] {
	user, err := u.userRepository.GetUserByEmail(request.Email)

	if err != nil {
		return model.NewErrorApiResponse[SignInResponse, string]("user", usererrors.UserNotFound.Error(), http.StatusNotFound)
	}

	isCorrect := user.IsPasswordCorrect(request.Password)

	if !isCorrect {
		return model.NewErrorApiResponse[SignInResponse, string]("email_or_password", usererrors.UserInvalidEmailOrPassword.Error(), http.StatusUnprocessableEntity)
	}

	tokenData := *model.NewTokenData(user.Id.Value(), user.Email.Value())
	accessToken, err := u.tokenizer.GenerateAccessToken(tokenData)

	if err != nil {
		return model.NewErrorApiResponse[SignInResponse, string]("accessToken", err.Error(), http.StatusInternalServerError)
	}

	refreshToken, err := u.tokenizer.GenerateRefreshToken(tokenData)

	if err != nil {
		return model.NewErrorApiResponse[SignInResponse, string]("refreshToken", err.Error(), http.StatusInternalServerError)
	}

	authUser := entity.NewAuthUser(entity.AuthUserInput{
		Id:           user.Id,
		Name:         user.Name,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

	response, err := NewResponse(SignInResponseInput{AuthUser: authUser})

	if err != nil {
		return model.NewErrorApiResponse[SignInResponse, string]("response", err.Error(), http.StatusInternalServerError)
	}

	return model.NewSuccessApiResponse[SignInResponse, string](response, http.StatusOK)
}
