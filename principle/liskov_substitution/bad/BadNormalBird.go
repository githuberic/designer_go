package bad

import "fmt"

// BadNormalBird实现了IBadBird接口
type BadNormalBird struct {
	iID   int
	sName string
}

func NewBadNormalBird(id int, name string) IBadBird {
	return &BadNormalBird{
		id,
		name,
	}
}
func (p *BadNormalBird) ID() int {
	return p.iID
}
func (p *BadNormalBird) Name() string {
	return p.sName
}
func (p *BadNormalBird) Tweet() error {
	fmt.Printf("BadNormalBird.Tweet, id=%v, name=%v\n", p.ID(), p.Name())
	return nil
}
func (p *BadNormalBird) Fly() error {
	fmt.Printf("BadNormalBird.Fly, id=%v, name=%v\n", p.ID(), p.Name())
	return nil
}
