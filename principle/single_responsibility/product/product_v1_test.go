package single_responsibility

import (
	"fmt"
	"testing"
	"designer_go/principle/single_responsibility/product/bad"
	good "designer_go/principle/single_responsibility/product/good"
)

func Test_SimpleResponsibility_Bad(t *testing.T) {
	product := bad.NewBadProduct("翔龙武士宝剑", 899)
	product.GetPrice(1)
	product.GetImage(1)
	fmt.Println(product.String())
}

func Test_SimpleResponsibility_Good(t *testing.T) {
	product := good.NewBadProduct("翔龙武士宝剑", 899)

	price := good.NewProductPrice(product)
	price.GetPrice(1)

	image := good.NewProductImage(product)
	image.GetImage(1)

	fmt.Println(product.String())
}
