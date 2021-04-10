package good

type S520 struct {
	product IProduct
}

func NewS520(ip IProduct) IScenario {
	return &SLive{product: ip}
}

func (p *S520) SetProduct(ip IProduct) {
	p.product = ip
}

func (p *S520) calc() float32 {
	return p.product.Price() * 0.8
}

func (p *S520) GetScenario() string {
	return "520大促场景"
}
