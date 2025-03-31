package signin

import (
	"github.com/rafaelbrunoss/general-server-go/internal/common/domain/model"
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service/logger"
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service/tokenizer"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/entity"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/repository"
)

type SignInUseCase struct {
	logger         logger.ILogger
	tokenizer      tokenizer.ITokenizer
	userRepository repository.IUserRepository
}

func NewSignInUseCase(
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

func (u *SignInUseCase) Execute(request *SignInCommand) (*SignInResponse, error) {
	user, err := u.userRepository.GetUserByEmail(request.Email)

	if err != nil {
		return nil, err
	}

	isCorrect := user.IsPasswordCorrect(request.Password)

	if !isCorrect {
		return nil, err
	}

	tokenData := *model.NewTokenData(user.Id.Value(), user.Email.Value())
	accessToken, err := u.tokenizer.GenerateAccessToken(tokenData)

	if err != nil {
		return nil, err
	}

	refreshToken, err := u.tokenizer.GenerateRefreshToken(tokenData)

	if err != nil {
		return nil, err
	}

	authUser := entity.NewAuthUser(entity.AuthUserInput{
		Id:           user.Id,
		Name:         user.Name,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

	response, err := NewSignInResponse(SignInResponseInput{AuthUser: authUser})

	if err != nil {
		return nil, err
	}

	return response, nil
}
