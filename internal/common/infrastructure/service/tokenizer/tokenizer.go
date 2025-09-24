package tokenizer

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	model "github.com/rafaelbrunotech/general-server-go/internal/common/domain/model"
)

var (
	UnexpectedSigningMethod = errors.New("unexpected signing method")
	CouldNotParseToken      = errors.New("could not parse the token")
	InvalidToken            = errors.New("invalid token")
	InvalidTokenClaims      = errors.New("invalid token claims")
)

type Tokenizer struct {
	secret []byte
}

func New() (*Tokenizer, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	return &Tokenizer{
		secret: secret,
	}, nil
}

func (t *Tokenizer) DecodeToken(token string) (*model.TokenData, error) {
	parsedToken, err := jwt.Parse(token, func(_token *jwt.Token) (interface{}, error) {
		_, ok := _token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, UnexpectedSigningMethod
		}

		return t.secret, nil
	})

	if err != nil {
		return nil, CouldNotParseToken
	}

	if !parsedToken.Valid {
		return nil, InvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, InvalidTokenClaims
	}

	userId := claims["userId"].(string)
	userEmail := claims["userEmail"].(string)

	if userId == "" || userEmail == "" {
		return nil, InvalidTokenClaims
	}

	tokenData := model.NewTokenData(userId, userEmail)

	return tokenData, nil
}

func (t *Tokenizer) VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(_token *jwt.Token) (interface{}, error) {
		_, ok := _token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, UnexpectedSigningMethod
		}

		return t.secret, nil
	})

	if err != nil {
		return CouldNotParseToken
	}

	if !parsedToken.Valid {
		return InvalidToken
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return InvalidTokenClaims
	}

	return nil
}

func (t *Tokenizer) GenerateAccessToken(tokenData model.TokenData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    tokenData.UserId,
		"userEmail": tokenData.UserEmail,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(t.secret)
}

func (t *Tokenizer) GenerateRefreshToken(tokenData model.TokenData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    tokenData.UserId,
		"userEmail": tokenData.UserEmail,
		"exp":       time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	return token.SignedString(t.secret)
}
