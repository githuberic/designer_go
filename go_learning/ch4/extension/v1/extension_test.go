package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// 类似集成 匿名类型
type Dog struct {
	Pet
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	// 调动父类的方法
	dog.Speak()
	dog.SpeakTo("Chao")
}
