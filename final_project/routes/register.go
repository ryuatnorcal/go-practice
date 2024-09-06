package routes

import (
	"net/http"
	"strconv"

	"example.com/final_project/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Registered for event"})

}

func cancelRegisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	_, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}
	var event models.Event

	err = event.Unregister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for event"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Cancelled registration for event"})
}
