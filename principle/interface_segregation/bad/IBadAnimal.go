package bad

/**
接口方法很多, 比较臃肿, 需要实现接口时负担很重
 */
type IBadAnimal interface {
	ID() int
	Name() string

	Eat() error
	Fly() error
	Swim() error
}
