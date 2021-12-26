package e1

import "fmt"

type WinImg struct {
}

func (p *WinImg) DoPaint(str string) {
	fmt.Println(str + " at winOS")
}
