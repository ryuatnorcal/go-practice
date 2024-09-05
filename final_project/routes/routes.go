package routes

// since getEvnts functions are in the same package, you don't need to import

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// you could pass anonymous function or regular function
	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent)

	server.POST("/events", createEvent)

	server.DELETE("/events/:id", deleteEvent)

	server.PUT("/events/:id", updateEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
