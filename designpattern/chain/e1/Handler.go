package e1

type Handler interface {
	Handle(content string)
	next(handler Handler, content string)
}
