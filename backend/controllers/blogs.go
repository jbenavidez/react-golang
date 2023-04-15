package controllers

import (
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type blog struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	CreateBy string `json:"create_by"`
}

func Blogs(c *gin.Context) {
	var albums = []blog{
		{Title: "1", Content: "Blue Train", CreateBy: "John Coltrane"},
	}
	fmt.Println(albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func CreateBlog(c *gin.Context) {
	var newblog blog

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newblog); err != nil {
		return
	}
	// Add the new album to the slice.
	fmt.Println("the new album", newblog.CreateBy)
	models.Create(newblog.Title, newblog.Content, newblog.CreateBy)
	c.IndentedJSON(http.StatusCreated, newblog)
}
