package main

import (
	"fmt"
	"log"

	"github.com/davidbolet/myTheresaTest/pkg/config"
	"github.com/davidbolet/myTheresaTest/pkg/controllers/products"
)

func main() {
	log.Printf("Starting myTheresa Promotions Test")
	appConfig := config.ParseProductsConfig()

	productsWebService := products.NewProductsWebService(appConfig)
	appPort := fmt.Sprintf(":%s", appConfig.AppPort)

	if err := productsWebService.Run(appPort); err != nil {
		log.Fatalf("Error starting HTTP server: " + err.Error())
	}
}
