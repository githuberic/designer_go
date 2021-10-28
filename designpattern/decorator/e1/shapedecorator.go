package e1

// 定义抽象装饰类
type ShapeDecorator struct {
	shape Shape
}
func (s *ShapeDecorator) ShapeDecorator(shape Shape) {
	s.shape = shape
}
func (s *ShapeDecorator) Draw() {
	s.shape.Draw()
}
