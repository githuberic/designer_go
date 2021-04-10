package v2

type IProduct interface {
	ID() int
	Name() string
	Price() float64
}

/**
拍品类, 实现IProduct接口
*/
type Product struct {
	iID    int
	sName  string
	fPrice float64
}

func NewProduct(id int, name string, price float64) *Product {
	return &Product{
		iID:    id,
		sName:  name,
		fPrice: price,
	}
}
func (p *Product) ID() int {
	return p.iID
}
func (p *Product) Name() string {
	return p.sName
}
func (p *Product) Price() float64 {
	return p.fPrice
}
