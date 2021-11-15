package mappers

import (
	"reflect"
	"testing"

	"github.com/davidbolet/myTheresaTest/pkg/model"
)

func TestMapPrice(t *testing.T) {
	type args struct {
		inPrice  int
		discount *model.Discount
	}
	percent := "30%"
	percent2 := "15%"
	percentError := "1%5%"
	tests := []struct {
		name    string
		args    args
		want    *model.Price
		wantErr bool
	}{
		{name: "Test Category", args: args{inPrice: 89000, discount: &model.Discount{Category: "boots", Percentage: percent}}, want: &model.Price{Original: 89000, Final: 62300, Discount: &percent, Currency: "EUR"}},
		{name: "Test SKU", args: args{inPrice: 10000, discount: &model.Discount{SkuList: &[]string{"000003"}, Percentage: percent2}}, want: &model.Price{Original: 10000, Final: 8500, Discount: &percent2, Currency: "EUR"}},
		{name: "Test no discount", args: args{inPrice: 89000}, want: &model.Price{Original: 89000, Final: 89000, Discount: nil, Currency: "EUR"}},
		{name: "Test error on discount", args: args{inPrice: 89000, discount: &model.Discount{Category: "boots", Percentage: percentError}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapPrice(tt.args.inPrice, tt.args.discount)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapProductToDto(t *testing.T) {
	type args struct {
		inProduct *model.Product
		discount  *model.Discount
	}
	percent := "30%"
	tests := []struct {
		name    string
		args    args
		want    *model.ProductDto
		wantErr bool
	}{
		{name: "Test Category", args: args{inProduct: &model.Product{Category: "boots", Price: 89000}, discount: &model.Discount{Category: "boots", Percentage: percent}}, want: &model.ProductDto{Category: "boots", Price: &model.Price{Original: 89000, Final: 62300, Discount: &percent, Currency: "EUR"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapProductToDto(tt.args.inProduct, tt.args.discount)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapProductToDto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapProductToDto() = %v, want %v", got, tt.want)
			}
		})
	}
}
