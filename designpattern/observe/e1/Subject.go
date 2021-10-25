package e1

// 抽象主题
type Subject interface {
	Add(o Observer)
	Send(str string)
}
