package good

import "fmt"

/**
GoodFlyableBird通过聚合GoodNormalBird实现IGoodBird接口, 通过提供Fly方法实现IFlyableBird子接口
*/
type GoodFlyableBird struct {
	GoodNormalBird
}

func NewGoodFlyableBird(id int, name string) IGoodBird {
	return &GoodFlyableBird{
		*NewGoodNormalBird(id, name),
	}
}
func (p *GoodFlyableBird) Fly() error {
	fmt.Printf("GoodFlyableBird.Fly, id=%v, name=%v\n", p.ID(), p.Name())
	return nil
}
