package e1

import "container/list"

type NewsSubject struct {
	title string
	list  *list.List
}

func (p *NewsSubject) Add(o Observer) {
	p.list.PushBack(o)
}
func (p *NewsSubject) Send(str string) {
	for i := p.list.Front(); i != nil; i = i.Next() {
		(i.Value).(Observer).Receive(p.title + " 发送的 " + str)
	}
}
