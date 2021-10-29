package e2

//使用 DrawAPI 接口创建抽象类 Shape。
type Shape struct {
	drawApi DrawApi
}

func (s *Shape) Shape(api DrawApi) {
	s.drawApi = api
}
