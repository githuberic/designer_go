package map_reduce

import (
	"fmt"
	"testing"
)

func Filter(arr []int, fn func(n int) bool) []int {
	var newArr = []int{}

	for _, it := range arr {
		if fn(it) {
			newArr = append(newArr, it)
		}
	}
	return newArr
}

func TestInvoke(t *testing.T) {
	var intSet = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	out := Filter(intSet, func(n int) bool {
		return n%2 == 1
	})
	fmt.Printf("%v\n",out)

	out = Filter(out, func(n int) bool {
		return n > 5
	})
	fmt.Printf("%v\n",out)
}