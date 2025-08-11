package registrations

import (
	"event-management-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	idStr := context.Param("id")
	userId := context.GetInt("userId")
	evtId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data", "error": err.Error()})
		return
	}

	event, err := models.GetEventById(evtId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not find event", "error": err.Error()})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register event", "error": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "successfully registered for event", "event": event})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}
