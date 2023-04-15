package controllers

import (
	"backend/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type blog struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	CreateBy string `json:"create_by"`
}

func Blogs(c *gin.Context) {
	var blogs = models.GeTAll()
	fmt.Println(blogs)
	c.IndentedJSON(http.StatusOK, blogs)
}

func CreateBlog(c *gin.Context) {
	var newblog blog

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newblog); err != nil {
		log.Println(err)
		fmt.Println("init request", err)
		return
	}
	// Add the new blog to the slice.
	models.Create(newblog.Title, newblog.Content, newblog.CreateBy)
	c.IndentedJSON(http.StatusCreated, newblog)
}

func BlogDetails(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	blog := models.GetBlog(id)
	c.IndentedJSON(http.StatusCreated, blog)
}

func BlogDelete(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	models.DeleteBlog(id)
	c.JSON(200, gin.H{"success": "Blog #" + idStr + " deleted"})

}
