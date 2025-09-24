package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbrunotech/general-server-go/internal/common/domain/model"
	"github.com/rafaelbrunotech/general-server-go/internal/packages/user/application"
	signin "github.com/rafaelbrunotech/general-server-go/internal/packages/user/application/command/sign-in"
	signup "github.com/rafaelbrunotech/general-server-go/internal/packages/user/application/command/sign-up"
)

type AuthController struct {
	useCases application.UseCases
}

func NewAuthController(useCases application.UseCases) *AuthController {
	return &AuthController{
		useCases: useCases,
	}
}

func (c *AuthController) SignIn(context *gin.Context) {
	var request signin.SignInCommandInput
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

	command, err := signin.NewCommand(request)

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

	response := c.useCases.Command.SignIn.Execute(command)

	context.JSON(response.Status, response)
}

func (c *AuthController) SignUp(context *gin.Context) {
	var request signup.SignUpCommandInput
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

	command, err := signup.NewCommand(request)

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

	response := c.useCases.Command.SignUp.Execute(command)

	context.JSON(response.Status, response)
}
