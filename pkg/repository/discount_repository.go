package repository

import "github.com/davidbolet/myTheresaTest/pkg/model"

type DiscountsRepository struct {
	discounts *[]model.Discount
}

//NewDiscountsRepository returns a discount
func NewDiscountsRepository() *DiscountsRepository {
	return &DiscountsRepository{discounts: &[]model.Discount{{ID: 1, Category: "boots", Percentage: "30%"}, {ID: 2, Percentage: "15%", SkuList: &[]string{"000003"}}}}
}

//ListDiscounts returns all discounts on database
func (r *DiscountsRepository) ListDiscounts() *[]model.Discount {
	return r.discounts
}

//Returns Discount that corresponds to category, nil if there's no discount
func (r *DiscountsRepository) FindByCategory(category string) *model.Discount {
	for _, v := range *r.discounts {
		if v.Category == category {
			return &v
		}
	}
	return nil
}

//Returns Discount that corresponds to given sku, nil if there's no discount
func (r *DiscountsRepository) FindBySku(sku string) *model.Discount {
	for _, v := range *r.discounts {
		if v.SkuList != nil {
			for _, s := range *v.SkuList {
				if s == sku {
					return &v
				}
			}
		}
	}
	return nil
}
