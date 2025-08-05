package main

import (
	"event-management-api/db"
	"event-management-api/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// Default() setup http server for us i. pre-configured
	server := gin.Default()
	// initalize db
	db.InitDB()

	// routes
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	// port: 8081
	server.Run(":8081")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBind(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	event.ID = 1
	event.USER_ID = 1
	event.SAVE()

	context.JSON(http.StatusAccepted, gin.H{
		"message": "Event created",
		"event":   event,
	})
}
