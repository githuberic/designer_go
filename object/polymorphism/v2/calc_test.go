package v2

import (
	"fmt"
	"testing"
)

// 定义函数类型
type Functype func(int, int) int

func Calc(x, y int, fun Functype) int {
	result := fun(x, y)
	return result
}

func Add(a, b int) int {
	return a + b
}

func Munic(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	return a * b
}

func TestVerify(t *testing.T)  {
	result := Calc(3,6,Add)
	fmt.Println(result)
}
