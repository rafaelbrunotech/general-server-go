package main

import (
	"os"

	app "github.com/rafaelbrunotech/general-server-go/internal"
	"github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	container := app.CreateContainer()

	server := gin.Default()
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	createAuthApi(server, container)
	createUsersApi(authenticated, container)

	port := os.Getenv("PORT")

	server.Run(":" + port)

	defer container.Common.DB.Client.Close()
}

func createAuthApi(server *gin.Engine, container *app.Container) {
	authController := container.User.AuthController

	server.POST("/sign-in", authController.SignIn)
	server.POST("/sign-up", authController.SignUp)
}

func createUsersApi(authenticated *gin.RouterGroup, container *app.Container) {
	userController := container.User.UserController

	authenticated.GET("/users", userController.GetUsers)
	authenticated.GET("/users/:id", userController.GetUserById)
	authenticated.PATCH("/users", userController.UpdateUser)
}
