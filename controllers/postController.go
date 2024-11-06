package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sunilsinghx/restapi-gin/initializers"
	"github.com/sunilsinghx/restapi-gin/models"
)

// POST
func CreatePost(c *gin.Context) {

	fmt.Println("Inside Create Post")
	//get data off from request
	var body models.Post
	c.Bind(&body)

	//create a post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {

		c.JSON(400, gin.H{
			"message": "Error storing in DB",
		})
		return
	}

	c.JSON(201, gin.H{"message": post})
	return

}

// GET
func GetAllPost(c *gin.Context) {
	var post []models.Post

	//get the post
	initializers.DB.Find(&post)

	//respond with result
	c.JSON(200, gin.H{
		"posts": post,
	})

}

// GET BY ID
func GetPostByID(c *gin.Context) {
	var post models.Post
	//get ID from url
	id := c.Param("id")

	//get the post
	initializers.DB.Find(&post, id)

	//respond with result
	c.JSON(200, gin.H{
		"posts": post,
	})
}

// PUT
func UpdatePost(c *gin.Context) {

	//get ID from request
	id := c.Param("id")

	//get Data from request
	var body models.Post
	c.Bind((&body))

	//find post from id
	var post models.Post
	initializers.DB.Find(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//send response
	c.JSON(204, gin.H{
		"Updated Post": post,
	})

}

// DELETE
func DeletePost(c *gin.Context) {

	//get id from request
	id := c.Param("id")

	//delete post from DB
	initializers.DB.Delete(&models.Post{}, id)

	//send response
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"message": "Post Deleted SuccessFully",
		"post":    posts,
	})

}
