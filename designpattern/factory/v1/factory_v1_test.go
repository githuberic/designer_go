package v1

import "testing"

func TestVerifyV1(t *testing.T) {
	var factory Factory

	red := factory.GetShape("Square")
	red.Draw()
}
