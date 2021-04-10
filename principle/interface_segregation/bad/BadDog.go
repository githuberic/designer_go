package bad

import (
	"errors"
	"fmt"
)

/**
BadDog实现IBadAnimal接口.
本来BadDog是不支持Fly()方法的, 但由于接口要求, 因此只能返回无意义错误.
*/
type BadDog struct {
	iID   int
	sName string
}

func NewBadDog(id int, name string) IBadAnimal {
	return &BadDog{
		iID:   id,
		sName: name,
	}
}

func (p *BadDog) ID() int {
	return p.iID
}
func (p *BadDog) Name() string {
	return p.sName
}
func (p *BadDog) Eat() error {
	fmt.Printf("%v/%v is eating\n", p.Name(), p.ID())
	return nil
}
func (p *BadDog) Fly() error {
	return errors.New(fmt.Sprintf("%v/%v cannot fly", p.Name(), p.ID()))
}
func (p *BadDog) Swim() error {
	fmt.Printf("%v/%v is swimming\n", p.Name(), p.ID())
	return nil
}
