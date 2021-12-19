package e2

import "fmt"

//创建实现了 DrawAPI 接口的实体桥接实现类
type GreenCircle struct {
}

func (s *GreenCircle) DrawCircle(radius int, x int, y int) {
	fmt.Println("radius、x、y:", radius, x, y)
}
