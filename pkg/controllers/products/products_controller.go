package products

import (
	"log"
	"net/http"

	"github.com/davidbolet/myTheresaTest/pkg/config"
	"github.com/davidbolet/myTheresaTest/pkg/mappers"
	"github.com/davidbolet/myTheresaTest/pkg/model"
	"github.com/davidbolet/myTheresaTest/pkg/repository"
	"github.com/gin-gonic/gin"
)

// ProductsController provides the REST api methods
type ProductsController struct {
	productsRepo   *repository.ProductsRepository
	discountsRespo *repository.DiscountsRepository
}

//NewProductsController returns a new instance of the controller
func NewProductsController(config config.ProductsConfig) *ProductsController {
	return &ProductsController{productsRepo: repository.NewProductsRepository(config.JsonPath), discountsRespo: repository.NewDiscountsRepository()}
}

//ListProducts lists user's assets
func (pc *ProductsController) ListProducts(ctx *gin.Context) {
	log.Printf("[START] Listing existing products")
	defer log.Printf("[FINISH] Listing products")
	result := []*model.ProductDto{}

	categoryFilter := ctx.Param("category")
	priceMin := ctx.Param("priceMin")
	priceMax := ctx.Param("priceMax")
	products := pc.productsRepo.ListProducts(&categoryFilter, &priceMin, &priceMax)
	for _, product := range *products {
		discount := pc.discountsRespo.FindBySku(product.Sku)
		if discount == nil {
			discount = pc.discountsRespo.FindByCategory(product.Category)
		}
		dto, err := mappers.MapProductToDto(&product, discount)
		if err != nil {
			ctx.AbortWithStatus(500)
			return
		}
		result = append(result, dto)
	}
	ctx.JSON(http.StatusOK, result)
}
