package good

type ProductPrice struct {
	*Product
}

func NewProductPrice(product *Product) *ProductPrice {
	return &ProductPrice{
		Product: product,
	}
}

func (p *ProductPrice) GetPrice(scenario int) {
	var price = p.price
	switch scenario {
	case normal:
	case P520:
		price = price * 0.52
	case P119:
		price = price * 0.9
	case PLive520:
		price = price * 0.5
	}
	p.Product.price = price
}
