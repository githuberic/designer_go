package good

type SOnePrice struct {
	product IProduct
}

func NewSOnePrice(ip IProduct) IScenario {
	return &SLive{product: ip}
}

func (p *SOnePrice) SetProduct(ip IProduct) {
	p.product = ip
}

func (p *SOnePrice) calc() float32 {
	return p.product.Price() * 0.8
}

func (p *SOnePrice) GetScenario() string {
	return "一口价场景"
}
