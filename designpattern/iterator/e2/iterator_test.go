package e2

import "testing"

func TestIterator(t *testing.T) {
	var aggregate Aggregate
	aggregate = NewNumbers(1, 10)

	IteratorPrint(aggregate.Iterator())
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}
