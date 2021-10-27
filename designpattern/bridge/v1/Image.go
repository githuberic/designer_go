package v1

type Image struct {
	drawImg DrawImg
}

func (p *Image) Image(drawImg DrawImg) {
	p.drawImg = drawImg
}
