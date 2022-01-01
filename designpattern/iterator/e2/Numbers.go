package e2

type Numbers struct {
	start, end int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

func (n *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		numbers: n,
		next:    n.start,
	}
}
