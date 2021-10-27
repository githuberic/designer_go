package v1

import "fmt"

type UnixImg struct {
}

func (p *UnixImg) DoPaint(str string) {
	fmt.Println(str + " at UnixOs")
}
