package e2

type Node interface {
	Interpret() int
}

func (n *ValNode) Interpret() int {
	return n.val
}
