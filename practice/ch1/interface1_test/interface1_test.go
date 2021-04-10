package interface1_test

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Sexual string
	Age int
}

func PrintPerson(p *Person) {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n", p.Name, p.Sexual, p.Age)
}

// 使用“成员函数”的方式叫“Receiver”，这种方式是一种封装，
// 因为 PrintPerson()本来就是和 Person强耦合的，
func (p *Person) Print() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n", p.Name, p.Sexual, p.Age)
}

func TestInterface1(t *testing.T)  {
	var p = Person{
		Name: "Hao Chen",
		Sexual: "Male",
		Age: 44,
	}

	PrintPerson(&p)
	p.Print()
}