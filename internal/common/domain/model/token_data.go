package model

type TokenData struct {
	UserId    string
	UserEmail string
}

func NewTokenData(userId string, userEmail string) *TokenData {
	t := &TokenData{
		UserId:    userId,
		UserEmail: userEmail,
	}

	return t
}
