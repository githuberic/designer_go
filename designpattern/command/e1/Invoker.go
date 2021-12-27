package e1

//调用者
type Invoker struct {
	command ICommand
}

func NewInvoker(command ICommand) *Invoker {
	invoker := &Invoker{}
	invoker.command = command
	return invoker
}
func (invoker *Invoker) Call() {
	if invoker.command != nil {
		invoker.command.Execute()
	}
}
