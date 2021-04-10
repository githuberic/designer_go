package product

// 结构体可以使用嵌套匿名结构体所有的字段和方法
// 首字母大写或者小写的字段、方法，都可以使用。
import (
	"fmt"
	"math/rand"
	"time"
)

/**
拍品
*/
type Product struct {
	Id    int
	Name  string
	Price float32
}

func NewProduct(name string, price float32) *Product {
	return &Product{Id: generate(), Name: name, Price: price}
}
func (p *Product) String() {
	fmt.Printf("Product,Id=%d,Name=%s,price=%v\n", p.Id, p.Name, p.Price)
}

/**
Id生成器
*/
func generate() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

/**
品牌馆-拍品
*/
type BrandProduct struct {
	Product
	brand string
}

type AuctionHouseProduct struct {
	Product
	AuctionHouse string
	EndTime      time.Time
}
