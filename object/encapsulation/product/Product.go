package product

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

/**
拍品
*/
type Product struct {
	id    int
	name  string
	price float32
}

/**
构造函数
*/
func NewProductV1(name string, price float32) *Product {
	t := new(Product)
	t.id = generate()
	t.name = name
	t.price = price
	return t
}

/**
构造函数
*/
func NewProductV2(name string, price float32) *Product {
	return &Product{id: generate(), name: name, price: price}
}

/**
Id生成器
*/
func generate() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

func (p *Product) SetName(name string) {
	p.name = name
}
func (p *Product) SetPrice(price float32) {
	p.price = price
}

func (p *Product) GetName() string {
	return p.name
}
func (p *Product) GetPrice() float32 {
	return p.price
}

func (p *Product) String() string {
	// 查看地址是否有被复制
	fmt.Printf("Address is %x", unsafe.Pointer(&p.name))
	return fmt.Sprintf("ID:%d/Name:%s/price:%f", p.id, p.name, p.price)
}

func (e Product) StringV2() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.name))
	return fmt.Sprintf("ID:%d-Name:%s-price:%f", e.id, e.name, e.price)
}
