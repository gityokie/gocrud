package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gityokie/gocrud/internal/controllers"
	"github.com/gityokie/gocrud/internal/store"
)

type DebugPrintRouteFunc func(httpMethod, absolutePath, handlerName string, nuHandlers int)

func SetupRouter(store store.Store, debugPrintRoute DebugPrintRouteFunc) *gin.Engine {
	r := gin.Default()

	r.LoadHTMLFiles("/home/maluki/disk2/gigs/gocrud/internal/routers/html/index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.POST("book", controllers.AddNewBook{}.Handle)

		v1.GET("book/:id", controllers.GetOneBook{}.Handle)

		v1.PUT("book/:id", controllers.PutOneBook{}.Handle)

		v1.DELETE("book/:id", controllers.DeleteBook{}.Handle)
	}

	return r
}
