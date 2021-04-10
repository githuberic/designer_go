package interface2_test

// 结构体嵌入
import (
	"fmt"
	"testing"
)

type WithName struct {
	Name string
}

type Country struct {
	WithName
}

type City struct {
	WithName
}

type Printable interface {
	PrintStr()
}

func (w WithName) PrintStr() {
	fmt.Println(w.Name)
}

func TestVerify(t *testing.T)  {
	c1 := Country {WithName{ "China"}}
	c2 := City { WithName{"Beijing"}}
	c1.PrintStr()
	c2.PrintStr()
}


