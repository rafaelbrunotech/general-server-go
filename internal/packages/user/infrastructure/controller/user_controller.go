package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbrunotech/general-server-go/internal/common/domain/model"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/application"
	updateuser "github.com/rafaelbrunotech/general-server-go/internal/packages/user/application/command/update-user"
	getuserbyid "github.com/rafaelbrunotech/general-server-go/internal/packages/user/application/query/get-user-by-id"
	getusers "github.com/rafaelbrunotech/general-server-go/internal/packages/user/application/query/get-users"
)

type UserController struct {
	useCases application.UseCases
}

func NewUserController(useCases application.UseCases) *UserController {
	return &UserController{
		useCases: useCases,
	}
}

func (c *UserController) GetUserById(context *gin.Context) {
	userId := context.Param("id")

	query, err := getuserbyid.NewQuery(getuserbyid.GetUserByIdQueryInput{UserId: userId})

	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			model.NewErrorApiResponse[any, string](
				"query",
				err.Error(),
				http.StatusUnprocessableEntity,
			),
		)
		return
	}

	response := c.useCases.Query.GetUserById.Execute(query)

	context.JSON(response.Status, response)
}

func (c *UserController) GetUsers(context *gin.Context) {
	query, err := getusers.NewQuery(getusers.GetUsersQueryInput{})

	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			model.NewErrorApiResponse[any, string](
				"query",
				err.Error(),
				http.StatusUnprocessableEntity,
			),
		)
		return
	}

	response := c.useCases.Query.GetUsers.Execute(query)

	context.JSON(response.Status, response)
}

func (c *UserController) UpdateUser(context *gin.Context) {
	var request updateuser.UpdateUserCommandInput
	err := context.ShouldBindJSON(&request)

	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			model.NewErrorApiResponse[any, string](
				"input",
				err.Error(),
				http.StatusUnprocessableEntity,
			),
		)
		return
	}

	command, err := updateuser.NewCommand(request)

	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			model.NewErrorApiResponse[any, string](
				"command",
				err.Error(),
				http.StatusUnprocessableEntity,
			),
		)
		return
	}

	response := c.useCases.Command.UpdateUser.Execute(command)

	context.JSON(response.Status, response)
}
