package main

import (
	"backend/controllers"
	"backend/middleware"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// init DB
	models.ConnectDB()
	models.DBMigrate()
	// ENDPOINT
	r := gin.New()
	r.Use(middleware.CORSMiddleware()) // set middleware
	r.GET("/blogs", controllers.Blogs)
	r.POST("/create", controllers.CreateBlog)
	r.GET("/blogs/:id", controllers.BlogDetails)
	//init project
	r.Run() // default port 8080
}
