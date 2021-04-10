package product

/**
直播
*/
type LiveProduct struct {
	Product
}

func NewLiveProduct(product Product) *LiveProduct {
	live := &LiveProduct{
		Product: product,
	}
	live.Product.CalcPrice = live.CalcPrice
	live.Product.SetKeyWord = live.SetKeyWord
	return live
}

func (p *LiveProduct) SetKeyWord() string {
	return "直播"
}

func (p *LiveProduct) CalcPrice() float32 {
	return p.price * 0.88
}
