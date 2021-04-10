package good

/**
实现IGoodAnimal接口, 提供动物的id,name等基本属性
*/
type GoodAnimalInfo struct {
	iID   int
	sName string
}

func (p *GoodAnimalInfo) ID() int {
	return p.iID
}
func (p *GoodAnimalInfo) Name() string {
	return p.sName
}
