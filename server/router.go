package server

import (
	"go-restful-product-search-webapp/controller"
	"go-restful-product-search-webapp/database"
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the gin router and routes.
func SetupRouter(db *database.DB) *gin.Engine {
	gin.SetMode("debug")
	r := gin.Default()
	r.Use(cors.Default())

	c := controller.New(db)

	v1 := r.Group("/api/v1/products")
	{
		v1.GET("search", c.SearchProducts)
		v1.GET("scan", c.ScanProducts)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "SHIPT",
		})
	})

	return r
}
