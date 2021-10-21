package v0_test

import (
	"fmt"
	"testing"
)

type Country struct {
	Name string
}
type City struct {
	Name string
}

type Printable interface {
	PrintStr()
}

func (c Country) PrintStr() {
	fmt.Println(c.Name)
}
func (c City) PrintStr() {
	fmt.Println(c.Name)
}

func TestInterface2(t *testing.T) {
	c1 := Country{"China"}
	c2 := City{"Beijing"}
	c1.PrintStr()
	c2.PrintStr()
}
