package interface3_test

// 接口完整性检查
// 检查一个对象是否实现了某接口所有的接口方法
import (
	"fmt"
	"testing"
)

type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s *Square) Sides() int {
	return 4
}

func TestVerify(t *testing.T) {
	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())

	/*
		校验是否全部实现一个接口
		var _ Shape = (*Square)(nil)
	*/
}
