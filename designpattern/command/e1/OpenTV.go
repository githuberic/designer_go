package e1

type OpenTV struct {
	TV *TV
}

func NewOpenTV() *OpenTV {
	commandTV := &OpenTV{}
	commandTV.TV = new(TV)
	return commandTV
}
func (command *OpenTV) Execute() {
	command.TV.Open()
}
