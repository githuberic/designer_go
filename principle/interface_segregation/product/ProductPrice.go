package product

type ProductPrice struct {
	*Product
}

func NewProductPrice(product *Product) IProduct {
	return &ProductPrice{
		Product: product,
	}
}

func (p *ProductPrice) GetPrice(scenario int) float32 {
	var price = p.price

	// 价格重新计算
	switch scenario {
	case normal:
	case P520:
		price = price * 0.52
	case P119:
		price = price * 0.9
	case PLive520:
		price = price * 0.5
	}
	return price
}
