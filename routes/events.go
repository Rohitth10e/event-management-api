package routes

import (
	"database/sql"
	"errors"
	"event-management-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	context.JSON(http.StatusOK, gin.H{"message": "sucecss", "events": events})
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

func getEvent(context *gin.Context) {
	idStr := context.Param("id")
	fmt.Println("idStr:", idStr)
	evtId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid event ID"})
		return
	}

	evt, err := models.GetEventById(evtId)
	if errors.Is(err, sql.ErrNoRows) {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to fetch event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Success", "event": evt})
}
