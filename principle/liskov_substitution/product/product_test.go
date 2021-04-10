package product

import (
	"fmt"
	"testing"
	"time"
)

func TestVerify(t *testing.T) {
	product := NewProduct("翔龙武士宝剑", 899.899)

	iPrice := NewBrandProductV2(product, "龙泉宝剑")
	if it, ok := iPrice.(IBrandProduct); ok {
		it.Brand()
	}

	image := NewFamousAuthorProductV2(product,"魏康大师")
	if it, ok := image.(IFamousAuthorProduct); ok {
		it.Author()
	}

	auctionHouse := NewPAuctionHouseV2(product,"西荣阁拍卖",time.Now())
	if it, ok := auctionHouse.(IAuctionHouseProduct); ok {
		it.GetInfo()
	}

	fmt.Println(image.String())
}
