package good

import "fmt"

/**
更好的Dog实现, 异味代码更少.
通过集成GoodAnimalInfo实现IGoodAnimal接口, 并选择性实现ISupportEat, ISupportSwim.
*/
type GoodDog struct {
	GoodAnimalInfo
}

func NewGoodDog(id int, name string) IGoodAnimal {
	return &GoodDog{
		GoodAnimalInfo{
			id,
			name,
		},
	}
}
func (p *GoodDog) Eat() error {
	fmt.Printf("%v/%v is eating\n", p.Name(), p.ID())
	return nil
}
func (p *GoodDog) Swim() error {
	fmt.Printf("%v/%v is swimming\n", p.Name(), p.ID())
	return nil
}
