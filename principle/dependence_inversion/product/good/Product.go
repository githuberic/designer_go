package good

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
func (p *Product) Price() float32 {
	return p.price
}

func (p *Product) CalcPrice(course IScenario) {
	course.SetProduct(p)
	// 价格重新计算
	p.price = course.calc()
}

func (p *Product) String() string {
	var buf bytes.Buffer
	buf.WriteString("Id=" + strconv.Itoa(p.iID))
	buf.WriteString(",Name=" + p.sName)
	buf.WriteString(",Price=" + fmt.Sprintf("%.2f", p.price))
	return buf.String()
}