package e2

import (
	"testing"
)

func TestBridge(t *testing.T) {
	shape := Shape{drawApi: new(RedCircle)}
	circle := Circle{shape: shape, x: 3, y: 6, radius: 2}
	circle.Draw()
}
