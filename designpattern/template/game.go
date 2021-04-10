package template

type Game struct {
	Initialize func()
	StartPlay  func()
	EndPlay    func()
}

//Play 模板基类的Play方法
func (p *Game) Play() {
	p.Initialize()
	p.StartPlay()
	p.EndPlay()
}
