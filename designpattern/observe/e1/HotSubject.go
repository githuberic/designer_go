package e1

import "container/list"

type HotSubject struct {
	title string
	list  *list.List
}

func (p *HotSubject) Add(o Observer) {
	p.list.PushBack(o)
}
func (p *HotSubject) Send(str string) {
	for i := p.list.Front(); i != nil; i = i.Next() {
		(i.Value).(Observer).Receive(p.title + " 发送的 " + str)
	}
}
