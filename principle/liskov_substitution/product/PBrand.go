package product

type PBrand struct {
	*Product
	brand string
}

func NewBrandProduct(id int, name string, price float32, brand string) IProduct {
	return &PBrand{
		&Product{iID: id, sName: name, price: price},
		brand,
	}
}

func NewBrandProductV2(product *Product, brand string) IProduct {
	return &PBrand{
		Product: product,
		brand:   brand,
	}
}

func (p *PBrand) Brand() string {
	return p.brand
}
