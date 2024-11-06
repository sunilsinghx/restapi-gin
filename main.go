package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sunilsinghx/restapi-gin/controllers"
	"github.com/sunilsinghx/restapi-gin/initializers"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {

	r := gin.Default()

	r.POST("/post", controllers.CreatePost)
	r.GET("/post", controllers.GetAllPost)
	r.GET("/post/:id", controllers.GetPostByID)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)

	r.Run()
}
