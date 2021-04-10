package good

type IScenario interface {
	SetProduct(IProduct)
	calc() float32
	GetScenario() string
}
