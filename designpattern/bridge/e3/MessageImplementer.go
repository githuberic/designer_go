package e3

type MessageImplementer interface {
	Send(text, to string)
}
