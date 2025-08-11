package events

import (
	"event-management-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	auth := server.Group("/")
	auth.Use(middlewares.Auth)
	auth.POST("/events", createEvent)
	auth.GET("/events/:id", getEvent)
	auth.PUT("/events/:id", updateEvent)
	auth.DELETE("/events/:id", deleteEvent)
}
