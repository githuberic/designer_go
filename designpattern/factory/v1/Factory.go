package v1

type AbstractFactory interface {
	GetShape(shapeType string) Shape
}

type Factory struct {
}

func (s *Factory) GetShape(shapeType string) Shape {
	if shapeType == "" {
		return nil
	}

	switch shapeType {
	case "Rectangle":
		return new(Rectangle)
	case "Square":
		return new(Square)
	default:
		return nil
	}
}
