package users

import (
	"event-management-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/signup", SignUp)
	auth := server.Group("/")
	auth.Use(middlewares.Auth)
	auth.GET("/users", GetAllUsers)
	server.POST("/login", Login)
}
