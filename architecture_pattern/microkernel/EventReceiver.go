package microkernel

type EventReceiver interface {
	OnEvent(evt Event)
}
