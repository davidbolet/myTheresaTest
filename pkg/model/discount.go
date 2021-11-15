package model

//Discount helds possible discount types
type Discount struct {
	ID         int
	Category   string
	SkuList    *[]string
	Percentage string
}
