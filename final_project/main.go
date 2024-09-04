package main

import (
	"net/http"

	"example.com/final_project/db"
	"example.com/final_project/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	// you could pass anonymous function or regular function
	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE

	server.POST("/events", createEvent)
	// run with port number with colon :
	server.Run(":8080")
}

// context is a structure that contains information about the request
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// bind json to struct
	// automatically map the request body to the event structure
	// make sure the request body is JSON and attributes are the same
	// if the incoming event is missing, just kept blank
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = 123
	event.UserId = 1
	event.Save()

	context.JSON(http.StatusCreated, event)

}
