package controllers

import (
	"backend/models"
	"fmt"
	"log"
	"net/http"

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
	// newAlbum.
	fmt.Println("init request")
	if err := c.BindJSON(&newblog); err != nil {
		log.Println(err)
		fmt.Println("init request", err)
		return
	}
	// Add the new album to the slice.
	fmt.Println("the new album", newblog.CreateBy)
	models.Create(newblog.Title, newblog.Content, newblog.CreateBy)
	c.IndentedJSON(http.StatusCreated, newblog)
}
