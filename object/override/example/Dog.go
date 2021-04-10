package example

import "fmt"

type AbstractDog struct {

	Sleep func()
}
func (p *AbstractDog) Eat() {
	fmt.Println("AbstractDog Eat")
	p.Sleep()
}
func (p *AbstractDog) Run() {
	fmt.Println("AbstractDog Run")
}


// 秋田犬
type Akita struct {
	AbstractDog
}
func NewAkita() *Akita {
	ptr := &Akita{}
	ptr.AbstractDog.Sleep = ptr.Sleep
	return ptr
}
func (p *Akita) Sleep() {
	fmt.Println("Akita Sleep")
}

// 拉布拉多犬
type Labrador struct {
	AbstractDog
}
func NewLabrador() *Labrador {
	ptr := &Labrador{}
	ptr.AbstractDog.Sleep = ptr.Sleep
	return ptr
}
func (p *Labrador) Sleep() {
	fmt.Println("Labrador Sleep")
}