package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbrunoss/general-server-go/internal/common/domain/model"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/application"
	signin "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/command/sign-in"
	signup "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/command/sign-up"
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
			model.NewErrorApiResponse(
				"",
				err.Error(),
				http.StatusUnprocessableEntity,
			),
		)
		return
	}

	command, err := signin.NewSignInCommand(request)

	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			model.NewErrorApiResponse(
				"",
				err.Error(),
				http.StatusUnprocessableEntity,
			),
		)
		return
	}

	response, err := c.useCases.Command.SignIn.Execute(command)

	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			model.NewErrorApiResponse(
				"",
				err.Error(),
				http.StatusUnprocessableEntity,
			),
		)
		return
	}

	context.JSON(http.StatusOK, model.NewSuccessApiResponse(response, http.StatusOK))
}

func (c *AuthController) SignUp(context *gin.Context) {
	var request signup.SignUpCommandInput
	err := context.ShouldBindJSON(&request)

	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			model.NewErrorApiResponse(
				"",
				err.Error(),
				http.StatusUnprocessableEntity,
			),
		)
		return
	}

	command, err := signup.NewSignUpCommand(request)

	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			model.NewErrorApiResponse(
				"",
				err.Error(),
				http.StatusUnprocessableEntity,
			),
		)
		return
	}

	response, err := c.useCases.Command.SignUp.Execute(command)

	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			model.NewErrorApiResponse(
				"",
				err.Error(),
				http.StatusUnprocessableEntity,
			),
		)
		return
	}

	context.JSON(http.StatusOK, model.NewSuccessApiResponse(response, http.StatusOK))
}
