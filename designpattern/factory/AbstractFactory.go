package factory

import (
	"fmt"
)

type Color interface {
	Fill()
}

type Red struct {
}

func (red *Red) Fill() {
	fmt.Println("Red fill")
}

type Green struct {
}

func (s *Green) Fill() {
	fmt.Println("Green Fill")
}

type Blue struct {
}

func (s *Blue) Fill() {
	fmt.Println("Blue Fill")
}

type AbstractFactory interface {
	GetColor(colorType string) Color
}

type AbsFactory struct {
}

func (*AbsFactory) GetColor(color string) Color {
	if color == "" {
		return nil
	}

	switch color {
	case "Red":
		return new(Red)
	case "Green":
		return new(Green)
	case "Blue":
		return new(Blue)
	default:
		return nil
	}
}
