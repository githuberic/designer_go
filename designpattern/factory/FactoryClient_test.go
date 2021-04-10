package factory

import "testing"

func TestVerifyV1(t *testing.T) {
	var factory AbsFactory

	red := factory.GetColor("Red")
	red.Fill()

	green := factory.GetColor("Green")
	green.Fill()

	blue := factory.GetColor("Blue")
	blue.Fill()
}
