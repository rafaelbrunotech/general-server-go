package tokenizer

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	model "github.com/rafaelbrunoss/general-server-go/internal/common/domain/model"
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
			return nil, errors.New("unexpected signing method")
		}

		return t.secret, nil
	})

	if err != nil {
		return nil, errors.New("could not parse the token")
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid token claims")
	}

	userId := claims["userId"].(string)
	// userEmail := claims["userEmail"].(string)
	tokenData := model.NewTokenData(userId, "")

	return tokenData, nil
}

func (t *Tokenizer) VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(_token *jwt.Token) (interface{}, error) {
		_, ok := _token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return t.secret, nil
	})

	if err != nil {
		return errors.New("could not parse the token")
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return errors.New("invalid token claims")
	}

	return nil
}

func (t *Tokenizer) GenerateAccessToken(tokenData model.TokenData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": tokenData.UserId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(t.secret)
}

func (t *Tokenizer) GenerateRefreshToken(tokenData model.TokenData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": tokenData.UserId,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	return token.SignedString(t.secret)
}
