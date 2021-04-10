package employee

import (
	"fmt"
	"testing"
)

func TestVerifyV1(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()

	person := Factory("tm")
	person.DoPBC()

	person = Factory("pm")
	person.DoPBC()
}
