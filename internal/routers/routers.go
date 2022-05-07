package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gityokie/gocrud/internal/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("book", controllers.ListBook)
		v1.POST("book", controllers.AddNewBook)
		v1.GET("book/:id", controllers.GetOneBook)
		v1.PUT("book/:id", controllers.PutOneBook)
		v1.DELETE("book/:id", controllers.DeleteBook)
	}

	return r
}
