package e2

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Printf("%#v\n", c)
	}
}
