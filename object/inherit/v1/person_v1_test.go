package v1

// 结构体可以使用嵌套匿名结构体所有的字段和方法
// 首字母大写或者小写的字段、方法，都可以使用。
import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	age  int
}
func (p *Person) speak() {
	fmt.Printf("I'm %s,age %v", p.Name, p.age)
}

type Programmer struct {
	Person
	Language string
}

func TestVerifyV1(t *testing.T)  {
	programmer := Programmer{
		Person : Person{
			Name: "lgq",
			age: 18,
		},
		Language: "golang",
	}
	// 匿名结构体字段访问可以简化
	programmer.speak()
	// 首字母大写或者小写的字段、方法，都可以使用。
	programmer.Person.speak()
}