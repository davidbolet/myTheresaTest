package config

import (
	"flag"
	"strconv"

	"github.com/davidbolet/myTheresaTest/pkg/util"
)

type ProductsConfig struct {
	AppPort          string
	TimeoutInSeconds int
	JsonPath         string
}

//ParseAssetConfig parses configuration
func ParseProductsConfig() ProductsConfig {
	config := ProductsConfig{}
	flag.StringVar(&config.AppPort, "p", util.GetEnv("PRODUCTS_PORT", "8088"), "server listen address")
	to, err := strconv.Atoi(util.GetEnv("PRODUCTS_TIMEOUT", "5"))
	if err != nil {
		panic("PRODUCTS_TIMEOUT given value is not a number")
	}
	flag.IntVar(&config.TimeoutInSeconds, "t", to, "Timeout for endpoint")
	flag.StringVar(&config.JsonPath, "db", util.GetEnv("JSONPATH", "products.json"), "json database file")
	flag.Parse()
	return config
}
