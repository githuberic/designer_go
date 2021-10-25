package e1

import "fmt"

type aObserver struct {
	name string
}

func (p *aObserver) Receive(str string) {
	fmt.Println("aObserver," + p.name + ",Received" + str)
}
