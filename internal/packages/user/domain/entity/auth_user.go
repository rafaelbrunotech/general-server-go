package entity

import (
	valueobject "github.com/rafaelbrunotech/general-server-go/internal/common/domain/value-object"
)

type AuthUserInput struct {
	Id           *valueobject.Id
	Name         string
	AccessToken  string
	RefreshToken string
}

type AuthUser struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewAuthUser(input AuthUserInput) AuthUser {
	return AuthUser{
		Id:           input.Id.Value(),
		Name:         input.Name,
		AccessToken:  input.AccessToken,
		RefreshToken: input.RefreshToken,
	}
}
