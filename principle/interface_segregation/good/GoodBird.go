package good

import "fmt"

/**
更好的Bird实现, 异味代码更少.
通过集成GoodAnimalInfo实现IGoodAnimal接口, 并选择性实现ISupportEat, ISupportFly.
*/
type GoodBird struct {
	GoodAnimalInfo
}

func NewGoodBird(id int, name string) IGoodAnimal {
	return &GoodBird{
		GoodAnimalInfo{
			id,
			name,
		},
	}
}

func (p *GoodBird) Eat() error {
	fmt.Printf("%v/%v is eating\n", p.Name(), p.ID())
	return nil
}
func (p *GoodBird) Fly() error {
	fmt.Printf("%v/%v is flying\n", p.Name(), p.ID())
	return nil
}
