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
}

func NewProduct(name string, price float32) *Product {
	rand.Seed(time.Now().UnixNano())
	return &Product{
		iID:   rand.Int(),
		sName: name,
		price: price,
	}
}

func (p *Product) ID() int {
	return p.iID
}
func (p *Product) Name() string {
	return p.sName
}

func (p *Product) String() string {
	var buf bytes.Buffer
	buf.WriteString("Id=" + strconv.Itoa(p.iID))
	buf.WriteString(",Name=" + p.sName)
	buf.WriteString(",Price=" + fmt.Sprintf("%.2f", p.price))
	return buf.String()
}
