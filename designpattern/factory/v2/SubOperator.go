package v2

type SubOperator struct {
}

func (p *SubOperator) Operate(x, y int) int {
	return x - y
}
