package registrations

import (
	"event-management-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRegistrationRoute(server *gin.Engine) {
	auth := server.Group("/")
	auth.Use(middlewares.Auth)
	auth.POST("/events/:id/register", registerForEvent)
	auth.DELETE("/events/:id/cancel", cancelRegistration)
}
