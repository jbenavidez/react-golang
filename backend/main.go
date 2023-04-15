package main

import (
	"backend/controllers"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// init DB
	models.ConnectDB()
	models.DBMigrate()
	// ENDPOINT
	r := gin.Default()
	r.GET("/blogs", controllers.Blogs)
	r.POST("/create", controllers.CreateBlog)
	//init project
	r.Run() // default port 8080
}
