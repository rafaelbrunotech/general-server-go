package service

import (
	model "github.com/rafaelbrunoss/general-server-go/internal/common/domain/model"
)

type ITokenizer interface {
	DecodeToken(token string) (*model.TokenData, error)
	GenerateAccessToken(tokenData model.TokenData) (string, error)
	GenerateRefreshToken(tokenData model.TokenData) (string, error)
	VerifyToken(token string) error
}
