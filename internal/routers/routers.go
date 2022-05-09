package routers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gityokie/gocrud/internal/controllers"
	"github.com/gityokie/gocrud/internal/logger"
	"github.com/gityokie/gocrud/internal/store"
)

type DebugPrintRouteFunc func(httpMethod, absolutePath, handlerName string, nuHandlers int)

func SetupRouter(store store.Store, debugPrintRoute DebugPrintRouteFunc) *gin.Engine {
	r := gin.Default()

	cwd, err := os.Getwd()
	if err != nil {
		logger.Log.Fatalf("failed to get current directory, err: %v", err)
	}

	r.LoadHTMLGlob(fmt.Sprintf("%s/internal/views/**/*", cwd))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
			"title": "Home website",
		})
	})

	r.GET("about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about/index.tmpl", gin.H{
			"title": "About website",
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
