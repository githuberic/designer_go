package v1

import "fmt"

type Shape interface {
	Draw()
}

type Rectangle struct {
}
func (s *Rectangle) Draw() {
	fmt.Print("draw Rectangle!")
}

type Square struct {
}
func (s *Square) Draw() {
	fmt.Println("draw Square!")
}
