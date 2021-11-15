package mappers

import (
	"strconv"
	"strings"

	"github.com/davidbolet/myTheresaTest/pkg/model"
)

//MapProductToDto maps a Product to the required format
func MapProductToDto(inProduct *model.Product, discount *model.Discount) (*model.ProductDto, error) {
	result := model.ProductDto{}
	result.Sku = inProduct.Sku
	result.Category = inProduct.Category
	result.Name = inProduct.Name

	priceRes, err := MapPrice(inProduct.Price, discount)
	if err != nil {
		return nil, err
	}
	result.Price = priceRes
	return &result, nil
}

//MapPrice maps the corresponding price, checking for format errors
func MapPrice(inPrice int, discount *model.Discount) (*model.Price, error) {
	result := model.Price{Currency: model.DEFAULT_CURRENCY, Original: inPrice, Final: inPrice}
	if discount != nil {
		//Check that discount field has a valid format
		discountNumStr := strings.Replace(*&discount.Percentage, "%", "", 1)
		discountNum, err := strconv.Atoi(discountNumStr)
		if err != nil {
			return nil, err
		}
		result.Final = inPrice - inPrice*discountNum/100
		result.Discount = &discount.Percentage
	}
	return &result, nil
}
