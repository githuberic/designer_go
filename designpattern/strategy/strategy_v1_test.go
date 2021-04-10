package strategy

import (
	"fmt"
	"testing"
)

type IComputer interface {
	Computer(x, y int) int
}

type Add struct {
}

func (p *Add) Computer(x, y int) int {
	return x + y
}

type Sub struct {
}

func (p *Sub) Computer(x, y int) int {
	return x - y
}

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

func TestVerify(t *testing.T)  {
	context := Context{A: 12,B: 3}
	context.SetContext(new(Add))
	fmt.Println(context.Result())

	context.SetContext(new(Sub))
	fmt.Println(context.Result())
}

// https://www.cnblogs.com/shi2310/p/11122155.html