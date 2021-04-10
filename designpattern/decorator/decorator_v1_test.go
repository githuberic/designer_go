package decorator

import (
	"testing"
)

func TestVerify(t *testing.T) {
	/*
	var rectangle = Rectangle{}
	rectangle.Draw()

	var circle = Circle{}
	circle.Draw()*/

	var shape Shape = new(Rectangle)

	// 增加Draw其他颜色属性
	/*
	var shapeDecorator ShapeDecorator
	shapeDecorator.shape = shape
	shapeDecorator.Draw()
	*/

	var redShapeDecorator = NewRedShapeDecorator(shape)
	redShapeDecorator.Draw()
}
