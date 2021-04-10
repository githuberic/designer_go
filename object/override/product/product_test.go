package product

import (
	"log"
	"testing"
)

func TestVerify(t *testing.T)  {
	product := NewProduct("翔龙武士宝剑", 899.899)
	liveP := NewJLProduct(*product)
	liveP.GetPrice()
	liveP.GetKeyWord()
	log.Print(liveP.String())
}

