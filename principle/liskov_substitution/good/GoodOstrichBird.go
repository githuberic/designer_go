package good

import "fmt"

/*
GoodOstrichBird通过聚合GoodNormalBird实现IGoodBird接口, 通过提供Run方法实现IRunnableBird子接口
*/
type GoodOstrichBird struct {
	GoodNormalBird
}

func NewGoodOstrichBird(id int, name string) IGoodBird {
	return &GoodOstrichBird{
		*NewGoodNormalBird(id, name),
	}
}
func (me *GoodOstrichBird) Run() error {
	fmt.Printf("GoodOstrichBird.Run, id=%v, name=%v\n", me.ID(), me.Name())
	return nil
}
