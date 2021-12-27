package e1

import "fmt"

type TV struct {
}

func NewTV() *TV {
	return &TV{}
}
func (tv *TV) ShutDown() error {
	fmt.Println("Close TV")
	return nil
}
func (tv *TV) Open() error {
	fmt.Println("Open TV")
	return nil
}
