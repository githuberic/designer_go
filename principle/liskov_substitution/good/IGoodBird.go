package good

/**
IGoodBird仅定义了最基本的方法集, 通过子接口IFlyableBird添加Fly方法, 通过子接口IRunnableBird添加Run方法
*/
type IGoodBird interface {
	ID() int
	Name() string

	// 是否可用鸣叫
	Tweet() error
}

type IFlyableBird interface {
	IGoodBird

	Fly() error
}

type IRunnableBird interface {
	IGoodBird

	Run() error
}
