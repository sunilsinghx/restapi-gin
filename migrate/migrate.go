package main

import (
	"github.com/sunilsinghx/restapi-gin/initializers"
	"github.com/sunilsinghx/restapi-gin/models"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	initializers.DB.AutoMigrate(&models.Post{})
}
