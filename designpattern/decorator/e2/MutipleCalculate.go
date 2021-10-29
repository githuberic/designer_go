package e2

//创建基于基准struct的乘法
type MutCalculate struct {
	Calculate
	num int
}

func NewMutCalculate(C Calculate, num int) *MutCalculate {
	return &MutCalculate{Calculate: C, num: num}
}

func (this *MutCalculate) Cal() int {
	return this.num * this.Calculate.Cal()
}
