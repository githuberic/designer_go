package product

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Product struct {
	iID   int
	sName string

	price float32
	CalcPrice func() float32

	keyword    string
	SetKeyWord func() string
}

func NewProduct(name string, price float32) *Product {
	rand.Seed(time.Now().UnixNano())
	return &Product{
		iID:   rand.Int(),
		sName: name,
		price: price,
		keyword: "放心拍文玩,微拍堂,值得信赖",
	}
}

func (p *Product) GetPrice() float32 {
	var price = p.price
	var scenarioPrice = p.CalcPrice()

	if price > scenarioPrice {
		price = scenarioPrice
	}
	p.price = price

	return price
}
func (p *Product) GetKeyWord() string {
	p.keyword = p.keyword + "," + p.SetKeyWord()
	return p.keyword
}

func (p *Product) String() string {
	var buf bytes.Buffer
	buf.WriteString("Id=" + strconv.Itoa(p.iID))
	buf.WriteString(",Name=" + p.sName)
	buf.WriteString(",Price=" + fmt.Sprintf("%.2f", p.price))
	buf.WriteString(",Keyword=" + p.keyword)
	return buf.String()
}
