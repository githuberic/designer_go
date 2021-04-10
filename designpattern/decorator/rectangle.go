package decorator

import "fmt"

type Rectangle struct {
}

func (p *Rectangle) NewRectangle() *Rectangle {
	return &Rectangle{}
}

func (r *Rectangle) Draw() {
	fmt.Println("Shape: Rectangle")
}
