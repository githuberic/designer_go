package e1

import (
	"fmt"
	"testing"
)

// Monkey结构体
type Monkey struct {
	Name string
}

// 声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (p *Monkey) climbing() {
	fmt.Println(p.Name, " 猴子生来会爬树")
}

// LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}

// 让LittleMonkey实现BirdAble
func (p *LittleMonkey) Flying() {
	fmt.Println(p.Name, "小猴子通过学习，会飞翔...")
}

//让LittleMonkey实现FishAble
func (p *LittleMonkey) Swimming() {
	fmt.Println(p.Name, "小猴子通过学习，会游泳..")
}

func TestVerifyV1(t *testing.T) {
	//创建一个LittleMonkey 实例
	monkey := LittleMonkey{
		Monkey {
			Name : "悟空",
		},
	}

	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}