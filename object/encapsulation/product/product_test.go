package product

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestVerify(t *testing.T) {
	t1 := NewProductV1("新疆和田玉", 18000.00)
	fmt.Printf("Address is %x\n", unsafe.Pointer(&t1.name))
	t.Log(t1.StringV2())
	fmt.Println(t1)
}
