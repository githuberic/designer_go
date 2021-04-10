package v2

import (
	"container/list"
	"fmt"
)

// 抽象观察者
type Observer interface {
	Receive(str string)
}

// 抽象主题
type Subject interface {
	Add(o Observer)
	Send(str string)
}

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

type aObserver struct {
	name string
}

func (p *aObserver) Receive(str string) {
	fmt.Println("aObserver," + p.name + ",Received" + str)
}

type bObserver struct {
	name string
}

func (p *bObserver) Receive(str string) {
	fmt.Println("bObserver," + p.name + ",Received" + str)
}
