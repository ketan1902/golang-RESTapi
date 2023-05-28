package migrate

import (
	"github.com/ketan1902/go-crud/initializers"
	"github.com/ketan1902/go-crud/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()

}

func main() {

	initializers.DB.AutoMigrate(&models.Post{})
}
