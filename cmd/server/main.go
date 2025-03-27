package main

import (
	app "github.com/rafaelbrunoss/general-server-go/internal"
	"github.com/rafaelbrunoss/general-server-go/internal/common/domain/config"

	"github.com/gin-gonic/gin"
)

func main() {
	container := app.CreateContainer()

	server := gin.Default()

	createUserApi(server, container)

	port := config.Env["PORT"]

	server.Run(":" + port)

	defer container.Common.DB.Client.Close()
}

func createUserApi(server *gin.Engine, container *app.Container) {
	userController := container.User.UserController

	server.GET("/users", userController.GetUsers)
	server.GET("/users/:id", userController.GetUserById)
	server.POST("/users", userController.CreateUser)
	server.PATCH("/users", userController.UpdateUser)
}
