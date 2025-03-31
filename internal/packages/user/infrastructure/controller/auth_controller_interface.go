package controller

import (
	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	SignIn(context *gin.Context)
	SignUp(context *gin.Context)
}
