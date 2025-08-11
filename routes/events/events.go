package events

import (
	"database/sql"
	"errors"
	"event-management-api/models"
	"event-management-api/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	context.JSON(http.StatusOK, gin.H{"message": "sucecss", "events": events})
}

func createEvent(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	isValid, userId := utils.VerifyToken(token)

	if !isValid {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var event models.Event
	err := context.ShouldBind(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	event.USER_ID = userId
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

func updateEvent(context *gin.Context) {
	idStr := context.Param("id")
	evtId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_, err = models.GetEventById(evtId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data", "error": err.Error()})
		return
	}

	updatedEvent.ID = evtId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event", "error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Success", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	idStr := context.Param("id")
	evtId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data", "error": err.Error()})
		return
	}

	_, err = models.GetEventById(evtId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
	}

	var deletedEvent models.Event
	deletedEvent.ID = evtId
	err = deletedEvent.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event", "error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Success", "event": deletedEvent})
}
