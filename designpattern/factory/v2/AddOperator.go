package v2

type AddOperator struct {
}
func (p *AddOperator) Operate(x, y int) int {
	return x + y
}
