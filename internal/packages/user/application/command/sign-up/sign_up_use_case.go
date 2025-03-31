package signup

import (
	"github.com/rafaelbrunoss/general-server-go/internal/common/domain/model"
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service/logger"
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service/tokenizer"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/entity"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/repository"
)

type SignUpUseCase struct {
	logger         logger.ILogger
	tokenizer      tokenizer.ITokenizer
	userRepository repository.IUserRepository
}

func NewSignUpUseCase(
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

func (u *SignUpUseCase) Execute(request *SignUpCommand) (*SignUpResponse, error) {
	user, err := entity.NewUser(
		entity.UserInput{
			Email:    request.Email,
			Name:     request.Name,
			Password: request.Password,
		},
	)

	if err != nil {
		return nil, err
	}

	err = u.userRepository.CreateUser(user)

	if err != nil {
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

	response, err := NewSignUpResponse(SignUpResponseInput{AuthUser: authUser})

	if err != nil {
		return nil, err
	}

	return response, nil
}
