package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/signup", SignUp)
	server.GET("/users", GetAllUsers)
	server.POST("/login", Login)
}
