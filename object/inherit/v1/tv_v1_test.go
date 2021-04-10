package v1

import (
	"fmt"
	"testing"
)

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

// 如一个 struct 嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现了多重继承。
type TV struct {
	Goods
	Brand
}

func TestVerify(t *testing.T) {
	// 嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	tv := TV{Goods{"电视机001", 5000.99}, Brand{"海尔", "山东"}}
	tv2 := TV{
		Goods{
			Price: 5000.99,
			Name:  "电视机002",
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
	}

	fmt.Println(tv)
	fmt.Println(tv2)
}
