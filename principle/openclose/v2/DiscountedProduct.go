package v2

/**
折扣接口
*/
type IDiscount interface {
	Discount() float64
}

/**
该结构体同时实现IProduct和IDiscount接口
*/
type DiscountedProduct struct {
	*Product
	fDiscount float64
}

func NewDiscountProduct(id int, name string, price float64, discount float64) *DiscountedProduct {
	return &DiscountedProduct{
		Product: &Product{
			iID:    id,
			sName:  name,
			fPrice: price,
		},
		fDiscount: discount,
	}
}

func NewDiscountProductV2(product *Product, discount float64) *DiscountedProduct {
	return &DiscountedProduct{
		Product:   product,
		fDiscount: discount,
	}
}

func (p *DiscountedProduct) Discount() float64 {
	return p.fDiscount
}
func (p *DiscountedProduct) Price() float64 {
	return p.fDiscount * p.Product.Price()
}
