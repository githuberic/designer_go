package good

import "fmt"

/**
GoodNormalBird提供对IGoodBird的基础实现
*/
type GoodNormalBird struct {
	iID   int
	sName string
}

func NewGoodNormalBird(id int, name string) *GoodNormalBird {
	return &GoodNormalBird{
		id,
		name,
	}
}
func (p *GoodNormalBird) ID() int {
	return p.iID
}
func (p *GoodNormalBird) Name() string {
	return p.sName
}
func (p *GoodNormalBird) Tweet() error {
	fmt.Printf("GoodNormalBird.Tweet, id=%v, name=%v\n", p.ID(), p.Name())
	return nil
}
