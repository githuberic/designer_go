package bad

import (
	"errors"
	"fmt"
)

/**
BadBird实现了IBadAnimal接口.
BadBird是不支持Swim()的, 但由于接口要求, 只能返回无意义的错误应付.
 */
type BadBird struct {
	iID   int
	sName string
}

func NewBadBird(id int, name string) IBadAnimal {
	return &BadBird{
		iID:   id,
		sName: name,
	}
}

func (p *BadBird) ID() int {
	return p.iID
}
func (p *BadBird) Name() string {
	return p.sName
}
func (p *BadBird) Eat() error {
	fmt.Printf("%v/%v is eating\n", p.Name(), p.ID())
	return nil
}
func (p *BadBird) Fly() error {
	fmt.Printf("%v/%v is flying\n", p.Name(), p.ID())
	return nil
}
func (p *BadBird) Swim() error {
	return errors.New(fmt.Sprintf("%v/%v cannot swimming", p.Name(), p.ID()))
}
