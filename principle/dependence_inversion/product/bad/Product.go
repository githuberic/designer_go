package bad

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/**
拍品
*/
type Product struct {
	id    int
	name  string
	price float32
}

func NewProduct(name string, price float32) *Product {
	rand.Seed(time.Now().UnixNano())
	return &Product{
		id:    rand.Int(),
		name:  name,
		price: price,
	}
}

func (p *Product) CalcLivePrice() {
	p.price = p.price * 0.88
}

func (p *Product) CalcOnePricePrice() {
	p.price = p.price * 0.80
}

func (p *Product) Calc520Price() {
	p.price = p.price * 0.67
}

func (p *Product) String() string {
	var buf bytes.Buffer
	buf.WriteString("Id=" + strconv.Itoa(p.id))
	buf.WriteString(",Name=" + p.name)
	buf.WriteString(",Price=" + fmt.Sprintf("%.2f", p.price))
	return buf.String()
}
