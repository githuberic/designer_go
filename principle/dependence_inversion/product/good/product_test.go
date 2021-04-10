package good

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T)  {
	product := NewProduct("翔龙武士宝剑", 899.899)
	product.CalcPrice(NewS520(product))
	fmt.Println(product.String())
}
