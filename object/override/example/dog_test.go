package example

import "testing"

func TestVerifyV2(t *testing.T)  {
	akita := NewAkita()
	akita.Eat()

	labrador := NewLabrador()
	labrador.Eat()
}
