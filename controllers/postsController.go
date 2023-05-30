package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ketan1902/go-crud/initializers"
	"github.com/ketan1902/go-crud/models"
)

func PostCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	posts := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&posts)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostIndex(c *gin.Context) {

	var post []models.Post

	initializers.DB.Find(&post)

	c.JSON(200, gin.H{
		"posts": post,
	})

}

func PostUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var posts models.Post
	initializers.DB.Find(&posts, id)

	initializers.DB.Model(&posts).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)

}
