package good

/**
拍品
*/
type IProduct interface {
	ID() int
	Name() string
	Price() float32

	CalcPrice(IScenario)
}
