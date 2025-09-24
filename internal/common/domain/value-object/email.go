package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

var (
    EmailCannotBeEmpty = errors.New("email cannot be empty")
    EmailInvalidFormat = errors.New("invalid email format")
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	email := strings.TrimSpace(value)

	if email == "" {
		return nil, EmailCannotBeEmpty
	}

	email = strings.ToLower(email)

	regex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	if !regex.MatchString(email) {
		return nil, EmailInvalidFormat
	}

	return &Email{value: email}, nil
}

func (e *Email) Value() string {
	return e.value
}
