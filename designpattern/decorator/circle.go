package decorator

import "fmt"

/**
Circle 圆形类
*/
type Circle struct {
}

func (p *Circle) NewCircle() *Circle {
	return &Circle{}
}

/**
Draw 输出方法，实现Shape接口
*/
func (p *Circle) Draw() {
	fmt.Println("Shape: Circle")
}
