package e2

type OriCalculate struct {
	num int
}

func NewOriCalculate(num int) *OriCalculate {
	return &OriCalculate{num: num}
}

func (o *OriCalculate) Cal() int {
	return o.num
}
