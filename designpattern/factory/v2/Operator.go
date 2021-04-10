package v2

type Operator interface {
	Operate(x, y int) int
}
