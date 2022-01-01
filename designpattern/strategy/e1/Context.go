package e1

var computer IComputer

type Context struct {
	A, B int
}

func (p *Context) SetContext(icomputer IComputer) {
	computer = icomputer
}
func (p *Context) Result() int {
	return computer.Computer(p.A, p.B)
}
