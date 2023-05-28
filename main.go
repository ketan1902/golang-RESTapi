package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ketan1902/go-crud/controllers"
	"github.com/ketan1902/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()

	initializers.ConnectToDB()

}

func main() {

	r := gin.Default()
	r.POST("/post", controllers.PostCreate)

	r.GET("/post", controllers.PostIndex)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()
}
