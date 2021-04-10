package bad

/**
不好的设计, 该接口未考虑某些鸟类是不能Fly的
*/
type IBadBird interface {
	ID() int
	Name() string

	Tweet() error
	Fly() error
}
