package e1

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	nameRepository := NewNameRepository()
	nameRepository.SetName("sy")
	nameRepository.SetName("lgq")
	nameRepository.SetName("ldx")
	nameRepository.SetName("lyy")

	it := nameRepository.GetIterator()
	for {
		name, ok := it()
		if !ok {
			break
		}
		fmt.Println("Get name: ", name)
	}
}
