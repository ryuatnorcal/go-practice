package main

import (
	"example.com/final_project/db"
	"example.com/final_project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	// run with port number with colon :
	server.Run(":8080")
}
