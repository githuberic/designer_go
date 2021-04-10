package bad

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	normal = iota
	P520
	P119
	PLive520
)

const (
	TengXun = iota
	Qiniu
)

type IProduct interface {
	ID() int
	Name() string

	GetPrice() float32
	GetImage(storage int) string
}

type Product struct {
	iID       int
	sName     string
	price     float64
	imagePath string
}

func NewBadProduct(name string, price float64) *Product {
	rand.Seed(time.Now().UnixNano())
	return &Product{
		iID:   rand.Int(),
		sName: name,
		price: price,
	}
}

func (p *Product) String() string {
	var buf bytes.Buffer
	buf.WriteString("Id=" + strconv.Itoa(p.iID))
	buf.WriteString(",Name=" + p.sName)
	buf.WriteString(",Price=" + fmt.Sprintf("%f", p.price))
	buf.WriteString(",Image path=" + p.imagePath)
	return buf.String()
}

func (p *Product) ID() int {
	return p.iID
}
func (p *Product) Name() string {
	return p.sName
}

func (p *Product) GetPrice(scenario int) {
	switch scenario {
	case normal:
	case P520:
		p.price = p.price * 0.52
	case P119:
		p.price = p.price * 0.9
	case PLive520:
		p.price = p.price * 0.5
	}
}

func (p *Product) GetImage(storage int) {
	extendPath := "jianlou.png"

	var path string
	switch storage {
	case TengXun:
		path = "https://www.tengxunyun.com/storage/wpt/image/live" + extendPath
	case Qiniu:
		path = "https://www.qiniuyunyun.com/storage/wpt/image/" + extendPath
	}
	p.imagePath = path
}
