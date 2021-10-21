package good

import (
	entity "designer_go/principle/composite_reuse/product/entity"
)

type ProductImage struct {
	ossConn IOSSConnection
}

func NewProductImage() *ProductImage {
	return &ProductImage{}
}

func (p *ProductImage) SetOSSConnection(ossConn IOSSConnection) {
	p.ossConn = ossConn
}

func (p *ProductImage) Upload(product *entity.Product) (error, int) {
	return p.ossConn.Upload("update image,", product.ID, product.Name, product.Price)
}

func (p *ProductImage) Download(it *entity.Product) (error, []byte) {
	return p.ossConn.Download("download image,", it.Name, it.Price, it.ID)
}
