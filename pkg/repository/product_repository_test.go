package repository

import (
	"testing"

	"github.com/davidbolet/myTheresaTest/pkg/model"
)

func Test_appliesFilters(t *testing.T) {
	type args struct {
		productFound   model.Product
		categoryFilter *string
		strPriceMin    *string
		strPriceMax    *string
	}
	categoryFilter1 := "boots"
	priceMinFilter := "11"
	priceMaxFilter := "20"
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Filter by category applies", args: args{productFound: model.Product{Category: "boots"}, categoryFilter: &categoryFilter1}, want: true},
		{name: "Filter by category not applies", args: args{productFound: model.Product{Category: "boots2"}, categoryFilter: &categoryFilter1}, want: false},
		{name: "Filter by price min not applies", args: args{productFound: model.Product{Category: "boots", Price: 10}, strPriceMin: &priceMinFilter}, want: false},
		{name: "Filter by price min applies", args: args{productFound: model.Product{Category: "boots", Price: 15}, strPriceMin: &priceMinFilter}, want: true},
		{name: "Filter by price max not applies", args: args{productFound: model.Product{Category: "boots", Price: 200}, strPriceMax: &priceMaxFilter}, want: false},
		{name: "Filter by price max applies", args: args{productFound: model.Product{Category: "boots", Price: 15}, strPriceMax: &priceMaxFilter}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appliesFilters(tt.args.productFound, tt.args.categoryFilter, tt.args.strPriceMin, tt.args.strPriceMax); got != tt.want {
				t.Errorf("appliesFilters() = %v, want %v", got, tt.want)
			}
		})
	}
}
