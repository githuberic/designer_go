package product

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	product := NewBadProduct("翔龙武士宝剑", 899)

	iPrice := NewProductPrice(product)
	if it, ok := iPrice.(IProductPrice); ok {
		product.price = it.GetPrice(1)
	}

	image := NewProductImage(product)
	if it, ok := image.(IProductImage); ok {
		product.imagePath = it.Image(1)
	}

	iKeyWord := NewPKeyWorld(product)
	if it, ok := iKeyWord.(IKeyWord); ok {
		product.keyword = it.Keyword()
	}

	fmt.Println(product.String())
}
