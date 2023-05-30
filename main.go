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
	r.POST("/posts", controllers.PostCreate)

	r.GET("/posts", controllers.PostIndex)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run()
}
