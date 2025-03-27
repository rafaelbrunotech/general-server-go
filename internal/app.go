package app

import (
	"github.com/rafaelbrunoss/general-server-go/internal/common"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user"
)

type Container struct {
	Common *common.Container
	User   *user.Container
}

func CreateContainer() *Container {
	commonContainer := common.NewContainer()

	userContainer := user.NewContainer(commonContainer)

	return &Container{
		Common: commonContainer,
		User:   userContainer,
	}
}
