package repository

import (
	"encoding/json"
	"fmt"
	"strconv"

	scribble "github.com/nanobox-io/golang-scribble"

	"github.com/davidbolet/myTheresaTest/pkg/model"
)

type ProductsRepository struct {
	db *scribble.Driver
}

func NewProductsRepository(path string) *ProductsRepository {
	db, err := scribble.New(path, nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	return &ProductsRepository{db: db}
}

func (p *ProductsRepository) ListProducts(categoryFilter *string, priceMin *string, priceMax *string) *[]model.Product {
	records, err := p.db.ReadAll("products")
	if err != nil {
		fmt.Println("Error", err)
	}
	products := []model.Product{}
	for _, f := range records {
		productFound := model.ProductList{}
		if err := json.Unmarshal([]byte(f), &productFound); err != nil {
			fmt.Println("Error", err)
		}
		for _, pr := range productFound.Products {
			if appliesFilters(pr, categoryFilter, priceMin, priceMax) {
				products = append(products, pr)
			}
		}
	}
	return &products
}

func appliesFilters(productFound model.Product, categoryFilter *string, strPriceMin *string, strPriceMax *string) bool {
	mustAdd := true
	if categoryFilter != nil && len(*categoryFilter) > 0 {
		mustAdd = mustAdd && productFound.Category == *categoryFilter
	}
	if mustAdd && strPriceMin != nil && len(*strPriceMin) > 0 {
		priceMin, err := strconv.Atoi(*strPriceMin)
		if err == nil {
			mustAdd = mustAdd && productFound.Price >= priceMin
		}
	}
	if mustAdd && strPriceMax != nil && len(*strPriceMax) > 0 {
		priceMax, err := strconv.Atoi(*strPriceMax)
		if err == nil {
			mustAdd = mustAdd && productFound.Price <= priceMax
		}
	}
	return mustAdd
}

func (p *ProductsRepository) AddProduct(product model.Product) bool {
	if err := p.db.Write("products", product.Sku, product); err != nil {
		fmt.Println("Error adding product", err)
		return false
	}
	return true
}
