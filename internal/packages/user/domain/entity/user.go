package entity

import (
	"time"

	valueobject "github.com/rafaelbrunotech/general-server-go/internal/common/domain/value-object"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Email    string
	Name     string
	Password string
}

type UserRestoreInput struct {
	Id        string     `json:"id"`
	Email     string     `json:"email"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	IsDeleted bool       `json:"isDeleted"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type User struct {
	Id        *valueobject.Id
	Email     *valueobject.Email
	Name      string
	Password  string
	CreatedAt time.Time
	DeletedAt *time.Time
	IsDeleted bool
	UpdatedAt time.Time
}

func NewUser(input UserInput) (*User, error) {
	email, err := valueobject.NewEmail(input.Email)

	if err != nil {
		return nil, err
	}

	user := &User{
		Id:        valueobject.NewId(),
		Email:     email,
		Name:      input.Name,
		CreatedAt: time.Now(),
		DeletedAt: nil,
		IsDeleted: false,
		UpdatedAt: time.Now(),
	}

	err = user.SetPassword(input.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) SetPassword(password string) error {
	securityLevel := 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), securityLevel)

	if err != nil {
		return err
	}

	u.Password = string(bytes)

	return nil
}

func (u *User) IsPasswordCorrect(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

func (u *User) SetEmail(email string) error {
	newEmail, err := valueobject.NewEmail(email)

	if err != nil {
		return err
	}

	u.Email = newEmail
	u.UpdatedAt = time.Now()

	return nil
}

func (u *User) SetName(name string) {
	u.Name = name
	u.UpdatedAt = time.Now()
}

func (u *User) Restore(input UserRestoreInput) error {
	email, err := valueobject.NewEmail(input.Email)

	if err != nil {
		return err
	}

	u.Id = valueobject.NewValue(input.Id)

	u.Email = email
	u.Name = input.Name
	u.Password = input.Password

	u.CreatedAt = input.CreatedAt
	u.DeletedAt = input.DeletedAt
	u.IsDeleted = input.IsDeleted
	u.UpdatedAt = input.UpdatedAt

	return nil
}
