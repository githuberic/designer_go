package e1

import "fmt"

type bObserver struct {
	name string
}

func (p *bObserver) Receive(str string) {
	fmt.Println("bObserver," + p.name + ",Received" + str)
}
