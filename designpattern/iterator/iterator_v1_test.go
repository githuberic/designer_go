package iterator

import (
	"fmt"
	"testing"
)

type NameRepository struct {
	names []string
}

func NewNameRepository() *NameRepository {
	return &NameRepository{
		names: make([]string, 0),
	}
}

func (p *NameRepository) GetIterator() func() (string, bool) {
	index := 0
	return func() (name string, ok bool) {
		if index >= len(p.names) {
			return
		}
		name, ok = p.names[index], true
		index++
		return
	}
}

func (p *NameRepository) SetName(name string) {
	p.names = append(p.names, name)
}

func TestVerify(t *testing.T) {
	nameRepository := NewNameRepository()
	nameRepository.SetName("sy")
	nameRepository.SetName("lgq")
	nameRepository.SetName("ldx")
	nameRepository.SetName("lyy")

	it := nameRepository.GetIterator()
	for{
		name, ok := it()
		if !ok{
			break
		}
		fmt.Println("Get name: ", name)
	}
}
