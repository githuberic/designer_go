package e1

import "fmt"

type BMPImage struct {
	image    Image
	radius   float32
	filename string
}

func NewBMPImage(image Image, fileName string, radius float32) *BMPImage {
	return &BMPImage{
		image:    image,
		filename: fileName,
		radius:   radius,
	}
}

func (p *BMPImage) Draw() {
	p.image.drawImg.DoPaint(p.filename)
	fmt.Printf("BMP Image %s,radius=%f", p.filename, p.radius)
}
