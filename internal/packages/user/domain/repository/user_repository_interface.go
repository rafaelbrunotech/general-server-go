package repository

import (
	valueobject "github.com/rafaelbrunoss/general-server-go/internal/common/domain/value-object"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/entity"
)

type IUserRepository interface {
	CreateUser(user *entity.User) error
	GetUserById(id *valueobject.Id) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	UpdateUser(user *entity.User) error
}
