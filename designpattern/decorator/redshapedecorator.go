package decorator

import "fmt"

// 定义具体的装饰类
type RedShapeDecorator struct {
	shapeDecorator Shape
}

func NewRedShapeDecorator(decoratedShape Shape) *RedShapeDecorator {
	return &RedShapeDecorator{shapeDecorator: decoratedShape}
}
func (s *RedShapeDecorator) Draw() {
	s.shapeDecorator.Draw()
	s.setRedBorder()
}
func (s *RedShapeDecorator) setRedBorder() {
	fmt.Println("Border Color: Red")
}
