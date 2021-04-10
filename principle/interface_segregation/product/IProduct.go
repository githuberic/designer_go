package product

/**
拍品
*/
type IProduct interface {
	ID() int
	Name() string
}

/**
拍品价格
*/
type IProductPrice interface {
	GetPrice(scenario int) float32
}

/**
拍品图片
*/
type IProductImage interface {
	Image(storage int) string
}

/**
拍品关键词
*/
type IKeyWord interface {
	Keyword() string
}

/**
品牌馆
*/
type IBrand interface {
	Brand() string
}

/**
名家
*/
type FamousAuthor interface {
	Author()
}
