package e2

// 创建基于基准struct的加法
type AddCalculate struct {
	Calculate
	num int
}

func NewAddCalculate(C Calculate, num int) *AddCalculate {
	return &AddCalculate{Calculate: C, num: num}
}

func (this *AddCalculate) Cal() int {
	return this.num + this.Calculate.Cal()
}
