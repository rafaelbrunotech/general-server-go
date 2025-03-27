package controller

import (
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	CreateUser(context *gin.Context)
	GetUserById(context *gin.Context)
	GetUsers(context *gin.Context)
	UpdateUser(context *gin.Context)
}
