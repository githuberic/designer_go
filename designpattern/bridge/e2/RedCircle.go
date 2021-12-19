package e2

import "fmt"

type RedCircle struct {
}

func (s *RedCircle) DrawCircle(radius int, x int, y int) {
	fmt.Println("radius、x、y:", radius, x, y)
}
