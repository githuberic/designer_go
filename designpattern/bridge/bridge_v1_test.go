package bridge

import (
	"fmt"
	"testing"
)

type DrawImg interface {
	DoPaint(str string)
}

type WinImg struct {
}

func (p *WinImg) DoPaint(str string) {
	fmt.Println(str + " at winOS")
}

type UnixImg struct {
}

func (p *UnixImg) DoPaint(str string) {
	fmt.Println(str + " at UnixOs")
}

type Image struct {
	drawImg DrawImg
}

func (p *Image) Image(drawImg DrawImg) {
	p.drawImg = drawImg
}

type BMPImage struct {
	image  Image
	radius string
}

func (p *BMPImage) BMPImage(image Image, radius string) {
	p.image = image
	p.radius = radius
}

func (p *BMPImage) Draw() {
	p.image.drawImg.DoPaint(p.radius)
	fmt.Printf("BMP Image %s", p.radius)
}

func TestVerify(t *testing.T) {
	var drawImg DrawImg = new(WinImg)

	var image = Image{drawImg: drawImg}
	var bmpImage = BMPImage{image: image, radius: "老刘"}
	bmpImage.Draw()
}
