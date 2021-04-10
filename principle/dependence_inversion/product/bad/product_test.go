package bad

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T)  {
	product := NewProduct("翔龙武士宝剑", 899.00)
	product.Calc520Price()
	fmt.Println(product.String())
	product.price = 899.00
	product.CalcOnePricePrice()
	fmt.Println(product.String())
}
