package entity

import (
	"time"

	valueobject "github.com/rafaelbrunoss/general-server-go/internal/common/domain/value-object"
)

type UserInput struct {
	Name string
}

type UserRestoreInput struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	Id        *valueobject.Id
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(input UserInput) (*User, error) {
	return &User{
		Id:        valueobject.NewId(),
		Name:      input.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (u *User) UpdateName(name string) {
	u.Name = name
	u.UpdatedAt = time.Now()
}

func (u *User) Restore(input UserRestoreInput) {
	u.Id = valueobject.NewValue(input.Id)
	u.Name = input.Name
	u.CreatedAt = input.CreatedAt
	u.UpdatedAt = input.UpdatedAt
}
