package good

type SLive struct {
	product IProduct
}

func NewSLive(ip IProduct) IScenario {
	return &SLive{product: ip}
}

func (p *SLive) SetProduct(ip IProduct) {
	p.product = ip
}

func (p *SLive) calc() float32 {
	return p.product.Price() * 0.88
}

func (p *SLive) GetScenario() string {
	return "直播场景"
}
