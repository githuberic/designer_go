package e2

//创建实现了 Shape 接口的实体类。
type Circle struct {
	shape  Shape
	x      int
	y      int
	radius int
}

func (s *Circle) Circle(x int, y int, radius int, drawAPI DrawApi) {
	s.shape.Shape(drawAPI)
	s.x = x
	s.y = y
	s.radius = radius
}

func (this *Circle) Draw() {
	this.shape.drawApi.DrawCircle(this.radius, this.x, this.y)
}
