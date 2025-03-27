package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/application"
	createuser "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/command/create-user"
	updateuser "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/command/update-user"
	getuserbyid "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/query/get-user-by-id"
	getusers "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/query/get-users"
)

type UserController struct {
	useCases application.UseCases
}

func NewUserController(useCases application.UseCases) *UserController {
	return &UserController{
		useCases: useCases,
	}
}

func (c *UserController) CreateUser(context *gin.Context) {
	var request createuser.CreateUserCommandInput
	err := context.ShouldBindJSON(&request)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	command, err := createuser.NewCreateUserCommand(request)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.useCases.Command.CreateUser.Execute(command)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user created"})
}

func (c *UserController) GetUserById(context *gin.Context) {
	userId := context.Param("id")

	query, err := getuserbyid.NewGetUserByIdQuery(getuserbyid.GetUserByIdQueryInput{UserId: userId})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.useCases.Query.GetUserById.Execute(query)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, response)
}

func (c *UserController) GetUsers(context *gin.Context) {
	query, err := getusers.NewGetUsersQuery(getusers.GetUsersQueryInput{})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.useCases.Query.GetUsers.Execute(query)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, response)
}

func (c *UserController) UpdateUser(context *gin.Context) {
	var request updateuser.UpdateUserCommandInput
	err := context.ShouldBindJSON(&request)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	command, err := updateuser.NewUpdateUserCommand(request)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.useCases.Command.UpdateUser.Execute(command)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user updated"})
}
