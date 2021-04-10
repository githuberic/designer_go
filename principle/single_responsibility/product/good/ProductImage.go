package good

type ProductImage struct {
	*Product
}

func NewProductImage(product *Product) *ProductImage {
	return &ProductImage{
		Product: product,
	}
}

func (p *ProductImage) GetImage(storage int) {
	extendPath := "jianlou.png"

	var path string
	switch storage {
	case TengXun:
		path = "https://www.tengxunyun.com/storage/wpt/image/live" + extendPath
	case Qiniu:
		path = "https://www.qiniuyunyun.com/storage/wpt/image/" + extendPath
	}
	p.imagePath = path
}
