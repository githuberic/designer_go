package interface2_test

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

type Stringable interface {
	ToString() string
}
func (c Country) ToString() string {
	return "Country = " + c.Name
}
func (c City) ToString() string{
	return "City = " + c.Name
}

func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}

func TestVerify(t *testing.T)  {
	d1 := Country {"USA"}
	d2 := City{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)
}


