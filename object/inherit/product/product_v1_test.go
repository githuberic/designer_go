package product

import (
	"testing"
	"time"
)

func TestVerifyV1(t *testing.T) {
	brandProduct := BrandProduct{
		Product: Product{Id: generate(), Name: "翔龙武士", Price: 988.00},
		brand:   "龙泉宝剑",
	}
	// 匿名结构体字段访问可以简化
	brandProduct.String()

	auctionHouseProduct := AuctionHouseProduct{
		Product:      Product{Id: generate(), Name: "印度老山檀香", Price: 5000.00},
		AuctionHouse: "世纪文博拍卖行",
		EndTime:      time.Now(),
	}
	// 匿名结构体字段访问可以简化
	auctionHouseProduct.String()
	t.Log(auctionHouseProduct)

	// 首字母大写或者小写的字段、方法，都可以使用。
	//programmer.Product.String()
}
