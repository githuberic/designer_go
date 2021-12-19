package v1

type Add struct {
}

func (p *Add) Computer(x, y int) int {
	return x + y
}
