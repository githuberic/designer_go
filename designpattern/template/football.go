package template

import "fmt"

type FootBall struct {
	Game
}

func NewFootBall() *FootBall {
	ft := new(FootBall)
	ft.Game.Initialize = ft.Initialize
	ft.Game.StartPlay = ft.StartPlay
	ft.Game.EndPlay = ft.EndPlay
	return ft
}

func (p *FootBall) Initialize()  {
	fmt.Println("Football game initialize")
}

//StartPlay 子类的StartPlay方法
func (p *FootBall) StartPlay() {
	fmt.Println("Football game started.")
}

//EndPlay 子类的EndPlay方法
func (p *FootBall) EndPlay() {
	fmt.Println("Football game Finished!")
}
