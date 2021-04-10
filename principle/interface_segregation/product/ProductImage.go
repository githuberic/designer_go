package product

type ProductImage struct {
	*Product
}

func NewProductImage(product *Product) IProduct {
	return &ProductImage{
		Product: product,
	}
}

func (p *ProductImage) GetImage(storage int) string {
	extendPath := p.imageName

	var path string
	switch storage {
	case TengXun:
		path = "https://www.tengxunyun.com/storage/wpt/image/live" + extendPath
	case Qiniu:
		path = "https://www.qiniuyunyun.com/storage/wpt/image/" + extendPath
	}
	return path
}
