package good

type IProduct interface {
	ID() int
	Name() string

	GetPrice() IProductPrice
	GetImage() IProductImage
}

type IProductPrice interface {
	GetPrice() float32
}
type IProductImage interface {
	GetImage(storage int) string
}
