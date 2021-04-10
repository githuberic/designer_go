package good

/**
将动物接口拆分为基本信息接口IGoodAnimal, 以及三个可选的能力接口:
ISupportEat, ISupportFly, ISupportSwim
*/
type IGoodAnimal interface {
	ID() int
	Name() string
}

type ISupportEat interface {
	Eat() error
}

type ISupportFly interface {
	Fly() error
}

type ISupportSwim interface {
	Swim() error
}
