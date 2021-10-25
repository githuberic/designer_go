package e1

// 抽象观察者
type Observer interface {
	Receive(str string)
}
