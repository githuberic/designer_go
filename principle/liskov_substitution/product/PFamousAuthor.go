package product

type PFamousAuthor struct {
	*Product
	author string
}

func NewFamousAuthorProduct(id int, name string, price float32, author string) IProduct {
	return &PFamousAuthor{
		&Product{iID: id, sName: name, price: price},
		author,
	}
}

func NewFamousAuthorProductV2(product *Product, author string) IProduct {
	return &PFamousAuthor{
		Product: product,
		author:  author,
	}
}

func (p *PFamousAuthor) Author() string {
	return p.author
}
