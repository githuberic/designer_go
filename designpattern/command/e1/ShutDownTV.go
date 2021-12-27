package e1

type ShutDownTV struct {
	TV *TV
}

func NewShutDownTV() *ShutDownTV {
	commandTV := &ShutDownTV{}
	commandTV.TV = NewTV()
	return commandTV
}
func (command *ShutDownTV) Execute() {
	command.TV.ShutDown()
}
