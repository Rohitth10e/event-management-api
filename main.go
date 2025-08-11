package main

import (
	"event-management-api/db"
	"event-management-api/routes/events"
	"event-management-api/routes/registrations"
	"event-management-api/routes/users"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default() setup http server for us i. pre-configured
	server := gin.Default()
	// initalize db
	db.InitDB()

	server.Use(func(c *gin.Context) {
		fmt.Println("Received request:", c.Request.Method, c.Request.URL.Path)
	})

	// event-routes
	events.RegisterEventRoutes(server)
	// user-routes
	users.RegisterUserRoutes(server)
	// registration-route
	registrations.RegisterRegistrationRoute(server)

	// port: 8081
	server.Run(":8081")
}
