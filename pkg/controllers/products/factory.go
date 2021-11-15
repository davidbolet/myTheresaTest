package products

import (
	"github.com/davidbolet/myTheresaTest/pkg/config"
	"github.com/gin-gonic/gin"
)

// RouterBuilder is the builder to create a GIN Router
type RouterBuilder struct {
	r *gin.Engine
}

// NewProductsWebService creates a new instance of the application
func NewProductsWebService(appConfig config.ProductsConfig) *gin.Engine {
	productsController := NewProductsController(appConfig)
	r := gin.New()
	r.GET("/products", productsController.ListProducts)
	return r
}
