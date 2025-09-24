package usererrors

import "errors"

var (
	UserNotFound               = errors.New("user not found")
	UserAlreadyExists          = errors.New("user already exists")
	UserInvalidEmailOrPassword = errors.New("user invalid email or password")
)
