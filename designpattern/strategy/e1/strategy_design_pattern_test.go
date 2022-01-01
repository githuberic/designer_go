package e1

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	context := Context{A: 12, B: 3}
	context.SetContext(new(Add))
	fmt.Println(context.Result())

	context.SetContext(new(Sub))
	fmt.Println(context.Result())
}
