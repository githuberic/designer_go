package chain

type Handler interface {
	Handle(content string)
	next(handler Handler, content string)
}
