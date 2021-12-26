package e1

import "fmt"

type PNGImage struct {
	image    Image
	radius   string
	fileName string
}

func NewPNGImage(image Image, fileName string, radius string) *PNGImage {
	return &PNGImage{image: image, fileName: fileName, radius: radius}
}

func (p *PNGImage) Draw() {
	p.image.drawImg.DoPaint(p.radius)
	fmt.Printf("PNG Image %s", p.radius)
}
