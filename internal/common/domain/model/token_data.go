package model

type TokenData struct {
	UserId string
}

func NewTokenData(userId string) *TokenData {
	t := &TokenData{
		UserId: userId,
	}

	return t
}
