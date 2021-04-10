package command

import (
	"fmt"
	"testing"
)

type TV struct {
}

func NewTV() *TV {
	return new(TV)
}
func (tv *TV) ShutDown() error {
	fmt.Println("Close TV")
	return nil
}
func (tv *TV) Open() error {
	fmt.Println("Open TV")
	return nil
}

//命令接口
type ICommand interface {
	Execute()
}

type ShutDownTV struct {
	TV *TV
}
func NewShutDownTV() *ShutDownTV {
	commandTV := new(ShutDownTV)
	commandTV.TV = NewTV()
	return commandTV
}
func (command *ShutDownTV) Execute()  {
	command.TV.ShutDown()
}


type OpenTV struct {
	TV *TV
}
func NewOpenTV() *OpenTV {
	commandTV := new(OpenTV)
	commandTV.TV = new(TV)
	return commandTV
}
func (command *OpenTV) Execute()  {
	command.TV.Open()
}



//调用者
type Invoker struct {
	command ICommand
}
func NewInvoker(command ICommand) *Invoker {
	invoker := new(Invoker)
	invoker.command = command
	return invoker
}
func (invoker *Invoker) Call()  {
	if invoker.command != nil{
		invoker.command.Execute()
	}
}

//命令模式，客户端通过调用者，传递不同的命令，然后不同的接受者对此进行处理
func TestVerify(t *testing.T) {
	invoker := NewInvoker(NewShutDownTV())
	invoker.Call()

	invokerV0 := NewInvoker(NewOpenTV())
	invokerV0.Call()
}

// from https://blog.csdn.net/m0_37681974/article/details/108117838