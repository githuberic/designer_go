package template

import "fmt"

type Basketball struct {
	Game
}

func NewBasketball() *Basketball {
	ft := new(Basketball)
	ft.Game.Initialize = ft.Initialize
	ft.Game.StartPlay = ft.StartPlay
	ft.Game.EndPlay = ft.EndPlay
	return ft
}

func (p *Basketball) Initialize()  {
	fmt.Println("Basketball game initialize")
}

//StartPlay 子类的StartPlay方法
func (p *Basketball) StartPlay() {
	fmt.Println("Basketball game started.")
}

//EndPlay 子类的EndPlay方法
func (p *Basketball) EndPlay() {
	fmt.Println("Basketball game Finished!")
}
