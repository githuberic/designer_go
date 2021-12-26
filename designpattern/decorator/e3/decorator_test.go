package e3

import (
	"fmt"
	"testing"
)

func TestDecorateE3(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	res := c.Calc()
	fmt.Printf("res %d\n", res)

	c = WarpMulDecorator(c, 8)
	res = c.Calc()

	fmt.Printf("res %d\n", res)
}
