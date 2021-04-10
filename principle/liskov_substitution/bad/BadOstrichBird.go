package bad

import (
	"errors"
	"fmt"
)

// BadOstrichBird通过继承BadNormalBird实现了IBadBird接口. 由于不支持Fly,
// 因此Fly方法抛出了错误. 额外添加了IBadBird未考虑到的Run方法. 该方法的调用要求调用方必须判断具体类型, 导致严重耦合.
type BadOstrichBird struct {
	BadNormalBird
}

func NewBadOstrichBird(id int, name string) IBadBird {
	return &BadOstrichBird{
		*(NewBadNormalBird(id, name).(*BadNormalBird)),
	}
}
func (p *BadOstrichBird) Fly() error {
	return errors.New(fmt.Sprintf("BadOstrichBird.Fly, cannot fly, id=%v, name=%v\n", p.ID(), p.Name()))
}
func (p *BadOstrichBird) Run() error {
	fmt.Printf("BadOstrichBird.Run, id=%v, name=%v\n", p.ID(), p.Name())
	return nil
}
