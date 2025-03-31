package controller

import (
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetUserById(context *gin.Context)
	GetUsers(context *gin.Context)
	UpdateUser(context *gin.Context)
}
