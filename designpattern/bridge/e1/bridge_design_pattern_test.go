package e1

import (
	"testing"
)

func TestVerify(t *testing.T) {
	//var drawImg DrawImg = new(WinImg)
	var image = Image{drawImg: new(WinImg)}
	var bmpImage = NewBMPImage(image, "a.bmp", 8.233)
	bmpImage.Draw()
}
