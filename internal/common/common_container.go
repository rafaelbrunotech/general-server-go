package common

import (
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/database"
	"github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/service"
)

type Services struct {
	Tokenizer service.ITokenizer
	Logger    service.ILogger
}

type Container struct {
	DB       *database.DB
	Services *Services
}

func NewContainer() *Container {
	services := provideServices()

	database := provideDatabase()

	return &Container{
		DB:       database,
		Services: services,
	}
}

func (c *Container) GetDB() *database.DB {
	return c.DB
}

func provideDatabase() *database.DB {
	database, err := database.InitDB()

	if err != nil {
		panic(err)
	}

	return database
}

func provideServices() *Services {
	logger, err := service.NewLogger()

	if err != nil {
		panic(err)
	}

	tokenizer, err := service.NewTokenizer()

	if err != nil {
		panic(err)
	}

	return &Services{
		Logger:    logger,
		Tokenizer: tokenizer,
	}
}
