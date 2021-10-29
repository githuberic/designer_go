package e1

type Target interface {
	Accept(Visitor)
}
