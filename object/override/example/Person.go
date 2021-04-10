package example

import "fmt"

type Person struct {
}
func (p *Person) Eat() {
	fmt.Println("Person Eat")
}
func (p *Person) Run() {
	fmt.Println("Person Run")
}
func (p *Person) Sleep() {
	fmt.Println("Person Sleep")
}


type Man struct {
	Person
}
func (p *Man) Eat() {
	// 类似于Java的 super.Eat()
	p.Person.Eat()
	fmt.Println("Man Eat")
}
func (p *Man) Run() {
	fmt.Println("Man Run")
}
