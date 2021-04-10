package v2

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	operator := NewOperatorFactory().CreateOperate("@")
	fmt.Printf("add result is %d\n", operator.Operate(1, 2)) //add result is 3
	return
}
