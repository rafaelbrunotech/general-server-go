package valueobject

import (
	"fmt"
	"regexp"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	email := strings.TrimSpace(value)

	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}

	email = strings.ToLower(email)

	regex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	if !regex.MatchString(email) {
		return nil, fmt.Errorf("invalid email format")
	}

	return &Email{value: email}, nil
}

func (e *Email) Value() string {
	return e.value
}
