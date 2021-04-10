package product

/**
捡漏
*/
type JLProduct struct {
	Product
}

func NewJLProduct(product Product) *JLProduct {
	jl := &JLProduct{
		Product: product,
	}
	jl.Product.CalcPrice = jl.CalcPrice
	jl.Product.SetKeyWord = jl.SetKeyWord
	return jl
}
func (p *JLProduct) SetKeyWord() string {
	return "捡漏,很划算"
}
func (p *JLProduct) CalcPrice() float32 {
	return p.price * 0.9
}
